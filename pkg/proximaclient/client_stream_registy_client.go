package proximaclient

import (
	"encoding/json"
	"fmt"
	http "github.com/hashicorp/go-retryablehttp"
	"github.com/proxima-one/stream-client-go/v2/pkg/connection"
	"io"
	"math"
	"math/rand"
	goHttp "net/http"
	"net/url"
	"time"
)

type StreamRegistryClient struct {
	client  *http.Client
	options StreamRegistryClientOptions
}

type StreamRegistryClientOptions struct {
	Endpoint        string
	RetryPolicy     connection.Policy
	DebugHttpOutput bool
}

type StreamFilter struct {
	Labels map[string]any `json:"labels"`
}

func NewStreamRegistryClient(options StreamRegistryClientOptions) *StreamRegistryClient {
	client := http.NewClient()
	// Exponential backoff https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/
	client.Backoff = func(min, max time.Duration, attemptNum int, resp *goHttp.Response) time.Duration {
		var backoff int64
		if attemptNum > 20 { // avoid overflow
			backoff = options.RetryPolicy.RetryMaxDelay.Milliseconds()
		} else {
			backoff = int64(math.Pow(2, float64(attemptNum))) * options.RetryPolicy.RetryBaseDelay.Milliseconds()
			if backoff > options.RetryPolicy.RetryMaxDelay.Milliseconds() {
				backoff = options.RetryPolicy.RetryMaxDelay.Milliseconds()
			}
		}
		println((time.Duration((backoff+int64(rand.Intn(int(backoff))))/2) * time.Millisecond).String())
		return time.Duration((backoff+int64(rand.Intn(int(backoff))))/2) * time.Millisecond
	}
	client.HTTPClient = &goHttp.Client{Timeout: options.RetryPolicy.Timeout}
	client.RetryMax = options.RetryPolicy.RetryCount
	if !options.DebugHttpOutput {
		client.Logger = nil
	}
	return &StreamRegistryClient{
		options: options,
		client:  client,
	}
}

func (client *StreamRegistryClient) GetStreamEndpoints(stream string, offset *Offset) ([]StreamEndpoint, error) {
	resp, err := client.client.Get(client.options.Endpoint + fmt.Sprintf("/streams/%s/offsets/%s/endpoints", stream, offset.String()))
	if err != nil {
		return nil, err
	}
	var res struct {
		Items []StreamEndpoint `json:"items"`
	}
	err = parseFromHttpResp(resp, &res)
	return res.Items, err
}

func (client *StreamRegistryClient) FindStream(stream string) (*Stream, error) {
	resp, err := client.client.Get(client.options.Endpoint + fmt.Sprintf("/streams/%s", stream))
	if err != nil {
		return nil, err
	}
	var res Stream
	err = parseFromHttpResp(resp, &res)
	return &res, err
}

func (client *StreamRegistryClient) FindStreams(filter *StreamFilter) ([]Stream, error) {
	postBody, _ := json.Marshal(filter)
	resp, err := client.client.Post(client.options.Endpoint+"/streams", "application/json", postBody)
	if err != nil {
		return nil, err
	}
	var res struct {
		Items []Stream `json:"items"`
	}
	err = parseFromHttpResp(resp, &res)
	return res.Items, err
}

func (client *StreamRegistryClient) GetStreams() ([]Stream, error) {
	resp, err := client.client.Get(client.options.Endpoint + "/streams")
	if err != nil {
		return nil, err
	}
	var res struct {
		Items []Stream `json:"items"`
	}
	err = parseFromHttpResp(resp, &res)
	return res.Items, err
}

func (client *StreamRegistryClient) FindOffset(stream string, height *int64, timestamp *time.Time) (*Offset, error) {
	if height == nil && timestamp == nil {
		return nil, fmt.Errorf("you should provide either height or timestamp")
	}

	queryParams, _ := url.ParseQuery("")
	if height != nil {
		queryParams.Add("height", fmt.Sprint(*height))
	}
	if timestamp != nil {
		queryParams.Add("timestamp", fmt.Sprint(timestamp.UnixMilli()))
	}

	resp, err := client.client.Get(client.options.Endpoint + fmt.Sprintf("/streams/%s/offsets/find?%s", stream, queryParams.Encode()))
	if err != nil {
		return nil, err
	}
	var res struct {
		Id string `json:"id"`
	}
	err = parseFromHttpResp(resp, &res)
	if err != nil {
		return nil, err
	}
	if res.Id == "" {
		return nil, OffsetNotFoundError()
	}
	return NewOffsetFromString(res.Id)
}

func parseFromHttpResp[T any](resp *goHttp.Response, obj T) error {
	text, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(text, obj)
	return err
}
