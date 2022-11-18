package model

type StreamStats struct {
	StartOffset Offset
	EndOffset   Offset
	Length      int64
	StorageSize int64
}
