package stream_model

type StreamEvent struct {
	Offset     Offset
	PrevOffset Offset
	Payload    []byte
	Timestamp  Timestamp
	Undo       bool
}
