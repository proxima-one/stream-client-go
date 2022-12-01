package proximaclient

type BufferedStreamReader struct {
	stream <-chan StreamEvent
}

func NewBufferedStreamReader(stream <-chan StreamEvent) *BufferedStreamReader {
	return &BufferedStreamReader{
		stream: stream,
	}
}

func (b *BufferedStreamReader) TryRead(maxBatchSize int) []StreamEvent {
	var res []StreamEvent
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
