package stream_registy

import "github.com/proxima-one/streamdb-client-go/pkg/model"

type SingleStreamDbRegistry struct {
	endpoint *model.StreamEndpoint
}

func NewSingleStreamDbRegistry(streamDbUri string) *SingleStreamDbRegistry {
	return &SingleStreamDbRegistry{endpoint: &model.StreamEndpoint{Uri: streamDbUri}}
}

func (registry *SingleStreamDbRegistry) GetStreamEndpoints(string, *model.Offset) ([]model.StreamEndpoint, error) {
	return []model.StreamEndpoint{*registry.endpoint}, nil
}
