package model

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
