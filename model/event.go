package model

import "time"

type Event struct {
	Payload   *[]byte
	Timestamp time.Time
	Undo      bool
}
