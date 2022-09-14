package model

import "time"

type Event struct {
	Payload   *[]byte
	Timestamp time.Time
	Undo      bool
}

func (e *Event) GetTimestamp() time.Time {
	return e.Timestamp
}

func (e *Event) GetPayload() *[]byte {
	return e.Payload
}

func (e *Event) GetUndo() bool {
	return e.Undo
}
