package proximaclient

type StreamMetadata struct {
	Description string         `json:"description"`
	Labels      map[string]any `json:"labels"`
}
