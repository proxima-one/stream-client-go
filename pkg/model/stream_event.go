package model

type StreamEvent struct {
	Offset    Offset
	Payload   []byte
	Timestamp Timestamp
	Undo      bool
}
