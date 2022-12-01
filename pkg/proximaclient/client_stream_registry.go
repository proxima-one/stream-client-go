package proximaclient

type StreamRegistry interface {
	GetStreamEndpoints(stream string, offset *Offset) ([]StreamEndpoint, error)
}
