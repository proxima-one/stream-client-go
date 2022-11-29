package stream_registy

import (
	"github.com/proxima-one/streamdb-client-go/pkg/stream_model"
)

type SingleStreamDbRegistry struct {
	endpoint *stream_model.StreamEndpoint
}

func NewSingleStreamDbRegistry(streamDbUri string) *SingleStreamDbRegistry {
	return &SingleStreamDbRegistry{endpoint: &stream_model.StreamEndpoint{Uri: streamDbUri}}
}

func (registry *SingleStreamDbRegistry) GetStreamEndpoints(string, *stream_model.Offset) ([]stream_model.StreamEndpoint, error) {
	return []stream_model.StreamEndpoint{*registry.endpoint}, nil
}
