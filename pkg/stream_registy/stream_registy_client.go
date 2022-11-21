package stream_registy

import (
	"encoding/json"
	"fmt"
	http "github.com/hashicorp/go-retryablehttp"
	"github.com/proxima-one/streamdb-client-go/pkg/connection"
	"github.com/proxima-one/streamdb-client-go/pkg/model"
	"io"
	"math"
	"math/rand"
	goHttp "net/http"
	"time"
)

type StreamRegistryClient struct {
	client  *http.Client
	options Options
}

type Options struct {
	Endpoint    string
	RetryPolicy connection.Policy
}

type StreamFilter struct {
	Labels map[string]string
}

func NewStreamRegistryClient(options Options) *StreamRegistryClient {
	client := http.NewClient()
	// Exponential backoff https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/
	client.Backoff = func(min, max time.Duration, attemptNum int, resp *goHttp.Response) time.Duration {
		var backoff int
		if attemptNum > 20 { // avoid overflow
			backoff = int(options.RetryPolicy.RetryMaxDelay)
		} else {
			backoff = int(math.Pow(2, float64(attemptNum))) * int(options.RetryPolicy.RetryBaseDelay)
			if backoff > int(options.RetryPolicy.RetryMaxDelay) {
				backoff = int(options.RetryPolicy.RetryMaxDelay)
			}
		}
		return time.Duration((backoff+rand.Intn(backoff))/2) * time.Millisecond
	}
	client.HTTPClient = &goHttp.Client{Timeout: options.RetryPolicy.Timeout}
	client.RetryMax = options.RetryPolicy.RetryCount
	return &StreamRegistryClient{
		options: options,
		client:  client,
	}
}

func (client *StreamRegistryClient) GetStreamEndpoints(stream string, offset *model.Offset) ([]model.StreamEndpoint, error) {
	resp, err := client.client.Get(client.options.Endpoint + fmt.Sprintf("/streams/%s/offsets/%s/endpoints", stream, offset.ToString()))
	if err != nil {
		return nil, err
	}
	var res struct {
		Items []model.StreamEndpoint `json:"items"`
	}
	err = parseFromHttpResp(resp, &res)
	return res.Items, err
}

func (client *StreamRegistryClient) FindStream(stream string) (*model.Stream, error) {
	resp, err := client.client.Get(client.options.Endpoint + fmt.Sprintf("/streams/%s", stream))
	if err != nil {
		return nil, err
	}
	var res model.Stream
	err = parseFromHttpResp(resp, &res)
	return &res, err
}

func (client *StreamRegistryClient) FindStreams(filter *StreamFilter) []model.Stream {
	return nil
}

func (client *StreamRegistryClient) GetStreams() []model.Stream {
	return nil
}

func (client *StreamRegistryClient) FindOffset(stream string, height *int64, timestamp *time.Time) *model.Offset {
	return nil
}

func parseFromHttpResp[T any](resp *goHttp.Response, obj T) error {
	text, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(text, obj)
	return err
}
