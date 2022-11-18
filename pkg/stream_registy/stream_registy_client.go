package stream_registy

import (
	http "github.com/hashicorp/go-retryablehttp"
	"github.com/proxima-one/streamdb-client-go/pkg/connection"
	"github.com/proxima-one/streamdb-client-go/pkg/model"
	"math"
	"math/rand"
	goHttp "net/http"
	"time"
)

type StreamRegistryClient struct {
	client http.Client
}

type Options struct {
	Endpoint    string
	RetryPolicy connection.Policy
}

type StreamFilter struct {
	Labels map[string]string
}

func NewStreamRegistryClient(options Options) *StreamRegistryClient {
	return &StreamRegistryClient{
		client: http.Client{
			HTTPClient: &goHttp.Client{Timeout: options.RetryPolicy.Timeout},
			RetryMax:   options.RetryPolicy.RetryCount,

			// Exponential backoff https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/
			Backoff: func(min, max time.Duration, attemptNum int, resp *goHttp.Response) time.Duration {
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
			},
		},
	}
}

func (client *StreamRegistryClient) GetStreamEndpoints(stream string, offset model.Offset) []model.StreamEndpoint {
	return nil
}

func (client *StreamRegistryClient) FindStream(stream string) *model.Stream {
	return nil
}

func (client *StreamRegistryClient) FindStreams(filter StreamFilter) []model.Stream {
	return nil
}

func (client *StreamRegistryClient) GetStreams() []model.Stream {
	return nil
}

func (client *StreamRegistryClient) FindOffset(stream string, height *int64, timestamp *time.Time) *model.Offset {
	return nil
}
