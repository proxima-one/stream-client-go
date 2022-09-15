package model

import (
	"encoding/json"
	"sync"
)

type TransitionPreprocessingResult struct {
	wg  sync.WaitGroup
	res any
	Err error
}

type TransitionPreprocessingFunc func(t *Transition) (any, error)

func NewTransitionPreprocessingResult(t *Transition, f TransitionPreprocessingFunc) *TransitionPreprocessingResult {
	res := &TransitionPreprocessingResult{
		wg:  sync.WaitGroup{},
		res: nil,
		Err: nil,
	}
	res.wg.Add(1)
	go func() {
		res.res, res.Err = f(t)
		res.wg.Done()
	}()
	return res
}

func (tr *TransitionPreprocessingResult) PreprocessingResult() (any, error) {
	tr.wg.Wait()
	return tr.res, tr.Err
}

func JsonParsingPreprocessFunc(transition *Transition) (any, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal(*transition.Event.Payload, &m)
	return m, err
}
