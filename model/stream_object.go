package model

import "time"

type ProximaStreamObject struct {
	Transition *Transition `json:"transition"`
	Preprocess *TransitionPreprocessingResult
}

func (o *ProximaStreamObject) GetTransition() *Transition {
	return o.Transition
}

func (o *ProximaStreamObject) GetPreprocess() *TransitionPreprocessingResult {
	return o.Preprocess
}

func (o *ProximaStreamObject) GetTimestamp() time.Time {
	return o.Transition.Event.Timestamp
}

func (o *ProximaStreamObject) GetUndo() bool {
	return o.Transition.Event.Undo
}
