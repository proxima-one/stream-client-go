package model

type StreamMetadata struct {
	Description string            `json:"description"`
	Labels      map[string]string `json:"labels"`
}
