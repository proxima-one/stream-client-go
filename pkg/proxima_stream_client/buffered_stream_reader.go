package proxima_stream_client

import (
	"github.com/proxima-one/streamdb-client-go/pkg/model"
)

type BufferedStreamReader struct {
	stream <-chan model.StreamEvent
}

func NewBufferedStreamReader(stream <-chan model.StreamEvent) *BufferedStreamReader {
	return &BufferedStreamReader{
		stream: stream,
	}
}

func (b *BufferedStreamReader) TryRead(maxBatchSize int) []model.StreamEvent {
	var res []model.StreamEvent
	event, ok := <-b.stream
	if !ok {
		return res
	}
	res = append(res, event)
	for i := 1; i < maxBatchSize; i++ {
		select {
		case event := <-b.stream:
			res = append(res, event)
		default:
			return res
		}
	}
	return res
}
