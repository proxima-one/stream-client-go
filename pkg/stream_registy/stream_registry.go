package stream_registy

import "github.com/proxima-one/streamdb-client-go/pkg/model"

type StreamRegistry interface {
	GetStreamEndpoints(stream string, offset *model.Offset) ([]model.StreamEndpoint, error)
}
