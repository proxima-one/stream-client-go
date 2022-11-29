package proxima_stream_client

import "github.com/proxima-one/streamdb-client-go/pkg/stream_model"

type BufferedStreamReader struct {
	stream <-chan stream_model.StreamEvent
}

func NewBufferedStreamReader(stream <-chan stream_model.StreamEvent) *BufferedStreamReader {
	return &BufferedStreamReader{
		stream: stream,
	}
}

func (b *BufferedStreamReader) TryRead(maxBatchSize int) []stream_model.StreamEvent {
	var res []stream_model.StreamEvent
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
