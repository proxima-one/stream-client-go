package stream_registy

import (
	"github.com/proxima-one/streamdb-client-go/pkg/stream_model"
)

type StreamRegistry interface {
	GetStreamEndpoints(stream string, offset *stream_model.Offset) ([]stream_model.StreamEndpoint, error)
}
