package model

type StreamStats struct {
	Id          string `json:"id"`
	StartOffset Offset `json:"start"`
	EndOffset   Offset `json:"end"`
	Length      int64  `json:"length"`
	StorageSize int64  `json:"storageSize"`
}
