package model

import (
	"context"
)

type Transition struct {
	NewState State
	Event    Event
	context  context.Context
}

func (t *Transition) GetNewState() State {
	return t.NewState
}

func (t *Transition) GetEvent() Event {
	return t.Event
}

func (t *Transition) IsEmpty() bool {
	return t.Event.Payload == nil
}
