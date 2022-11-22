package model

type StreamStats struct {
	Id          string  `json:"id"`
	Length      int64   `json:"length"`
	StartOffset *Offset `json:"start"`
	EndOffset   *Offset `json:"end"`
	StorageSize *int64  `json:"storageSize"`
}
