package stream_registy

import "github.com/proxima-one/streamdb-client-go/pkg/model"

type SingleStreamDbRegistry struct {
	endpoint *model.StreamEndpoint
}

func NewSingleStreamDbRegistry(endpoint *model.StreamEndpoint) *SingleStreamDbRegistry {
	return &SingleStreamDbRegistry{endpoint: endpoint}
}

func (registry *SingleStreamDbRegistry) GetStreamEndpoints(string, *model.Offset) ([]model.StreamEndpoint, error) {
	return []model.StreamEndpoint{*registry.endpoint}, nil
}
