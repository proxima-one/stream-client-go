package model

import (
	"context"
	"sync"
)

type Transition struct {
	NewState State
	Event    Event
	context  context.Context
}

//todo: add error group instead is empty hack
func (t Transition) IsEmpty() bool {
	return t.Event.Payload == nil
}

type TransitionPreprocessingResult struct {
	wg  sync.WaitGroup
	res any
	err error
}

type TransitionPreprocessingFunc func(t *Transition) (any, error)

func NewTransitionPreprocessingResult(t *Transition, f TransitionPreprocessingFunc) *TransitionPreprocessingResult {
	res := &TransitionPreprocessingResult{
		wg:  sync.WaitGroup{},
		res: nil,
		err: nil,
	}
	res.wg.Add(1)
	go func() {
		res.res, res.err = f(t)
		res.wg.Done()
	}()
	return res
}

func (tr *TransitionPreprocessingResult) PreprocessingResult() (any, error) {
	tr.wg.Wait()
	return tr.res, tr.err
}
