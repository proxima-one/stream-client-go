package model

type Transition struct {
	NewState State
	Event    Event
}

//todo: add error group instead is empty hack
func (t Transition) IsEmpty() bool {
	return t.Event.Payload == nil
}
