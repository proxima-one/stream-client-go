package proximaclient

type Stream struct {
	Name      string                    `json:"name"`
	Metadata  StreamMetadata            `json:"metadata"`
	Endpoints map[string]StreamEndpoint `json:"endpoints"`
}
