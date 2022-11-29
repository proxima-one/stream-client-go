package stream_model

type StreamEndpoint struct {
	Uri   string      `json:"uri"`
	Stats StreamStats `json:"stats"`
}
