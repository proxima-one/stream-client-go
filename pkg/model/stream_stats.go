package model

type StreamStats struct {
	StartOffset Offset `json:"start"`
	EndOffset   Offset `json:"end"`
	Length      int64  `json:"length"`
	StorageSize int64  `json:"storageSize"`
}
