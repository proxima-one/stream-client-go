package client

import (
	"context"
	"encoding/json"
	"github.com/proxima-one/streamdb-client-go/model"
	"sync"
)

type TransitionPreprocessingResult struct {
	wg  sync.WaitGroup
	res any
	err error
}

type TransitionPreprocessingFunc func(t *model.Transition) (any, error)

func NewTransitionPreprocessingResult(t *model.Transition, f TransitionPreprocessingFunc) *TransitionPreprocessingResult {
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

type ProximaStreamObject struct {
	Transition *model.Transition `json:"transition"`
	Preprocess *TransitionPreprocessingResult
}

func JsonParsingPreprocessFunc(transition *model.Transition) (any, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal(*transition.Event.Payload, &m)
	return m, err
}

type ProximaStreamSimpleReaderInterface interface {
	Start(ctx context.Context, option *StreamConnectionOption) error
	ReadNext() (*ProximaStreamObject, error)
}
