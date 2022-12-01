package proximaclient

type SingleStreamDbRegistry struct {
	endpoint *StreamEndpoint
}

func NewSingleStreamDbRegistry(streamDbUri string) *SingleStreamDbRegistry {
	return &SingleStreamDbRegistry{endpoint: &StreamEndpoint{Uri: streamDbUri}}
}

func (registry *SingleStreamDbRegistry) GetStreamEndpoints(string, *Offset) ([]StreamEndpoint, error) {
	return []StreamEndpoint{*registry.endpoint}, nil
}
