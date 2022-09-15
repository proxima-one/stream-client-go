package client

import (
	"context"
	"github.com/proxima-one/streamdb-client-go/config"
	"github.com/proxima-one/streamdb-client-go/model"
	"log"
	"sync"
	"time"
	//"google.golang.org/grpc/credentials" todo: add ssl and auth
)

const StreamConnectionOptionTypeRpc = "StreamConnectionOptionTypeRpc"
const StreamConnectionOptionTypeStream = "StreamConnectionOptionTypeStream"
const StreamConnectionOptionDefaultReconnectTime = time.Hour
const StreamConnectionOptionDefaultWatchdogTimeout = time.Second * 4

type StreamConnectionOption struct {
	Type            string
	ReconnectTime   time.Duration
	WatchDogTimeout time.Duration
}

type ProximaStreamSimpleReaderInterface interface {
	Start(ctx context.Context, option *StreamConnectionOption) error
	ReadNext() (*model.ProximaStreamObject, error)
}

func NewDefaultStreamConnectionOption() *StreamConnectionOption {
	return &StreamConnectionOption{
		Type:            StreamConnectionOptionTypeStream,
		ReconnectTime:   StreamConnectionOptionDefaultReconnectTime,
		WatchDogTimeout: StreamConnectionOptionDefaultWatchdogTimeout,
	}
}

func (conn *StreamConnectionOption) IsEmpty() bool {
	return conn.Type == ""
}

func (conn *StreamConnectionOption) WithReconnectTime(reconnectTime time.Duration) *StreamConnectionOption {
	conn.ReconnectTime = reconnectTime
	return conn
}

func (conn *StreamConnectionOption) WithWatchDogTimeout(watchDogTimeout time.Duration) *StreamConnectionOption {
	conn.WatchDogTimeout = watchDogTimeout
	return conn
}

func (conn *StreamConnectionOption) WithType(typeName string) *StreamConnectionOption {
	conn.Type = typeName
	return conn
}

func (conn *StreamConnectionOption) IsRpcBased() bool {
	return conn.Type == StreamConnectionOptionTypeRpc
}

func (conn *StreamConnectionOption) IsStreamBased() bool {
	return conn.Type == StreamConnectionOptionTypeStream
}

func NewStreamConnectionOption(typeName string) *StreamConnectionOption {
	res := NewDefaultStreamConnectionOption()
	return res.WithType(typeName)
}

func NewRpcBasedStreamConnectionOption() *StreamConnectionOption {
	return NewStreamConnectionOption(StreamConnectionOptionTypeRpc)
}

func NewStreamBasedStreamConnectionOption(reconnectTime time.Duration) *StreamConnectionOption {
	return NewStreamConnectionOption(StreamConnectionOptionTypeStream).WithReconnectTime(reconnectTime)
}

type StreamReader struct {
	config    config.Config
	client    *ProximaClient
	lastState model.State

	preprocessFunc model.TransitionPreprocessingFunc

	processedCount int

	streamConnectionOption StreamConnectionOption
	input                  *<-chan *model.Transition
	inputErr               *<-chan error

	output    chan *model.ProximaStreamObject
	outputErr chan error
	outputCtx context.Context

	isConnected     bool
	isConnectedLock sync.RWMutex

	mutex          *sync.Mutex
	watchdogCancel context.CancelFunc
	restartSync    chan struct{}
	stopSync       chan struct{}
}

func NewStreamReader(config config.Config, preprocess model.TransitionPreprocessingFunc) *StreamReader {
	res := &StreamReader{
		config:         config,
		lastState:      config.GetState(),
		preprocessFunc: preprocess,
		mutex:          &sync.Mutex{},
	}
	res.client = NewProximaClient(res.config)
	res.restartSync = make(chan struct{})
	res.stopSync = make(chan struct{})
	return res
}

func (reader *StreamReader) GetBufferLoadPercent() int {
	if reader.output == nil {
		return 0
	}
	load := float32(len(reader.output)) / float32(reader.config.GetChannelSize()) * 100
	return int(load)
}

func (reader *StreamReader) Metrics() (int, int) {
	return reader.processedCount, reader.GetBufferLoadPercent()
}

func (reader *StreamReader) IsConnected() bool {
	reader.isConnectedLock.RLock()
	defer reader.isConnectedLock.RUnlock()
	return reader.isConnected
}

func (reader *StreamReader) setIsConnected(val bool) {
	reader.isConnectedLock.Lock()
	reader.isConnected = val
	reader.isConnectedLock.Unlock()
}

func (reader *StreamReader) GetStreamID() string {
	return reader.config.GetStreamID()
}

func (reader *StreamReader) pushObject(obj *model.ProximaStreamObject) {
	if reader.output == nil {
		panic("output channel is nil")
	}
	reader.output <- obj
	reader.lastState = obj.Transition.NewState
	reader.processedCount++
	//here can make reconnection
}

func (reader *StreamReader) streamObjForTransition(transition *model.Transition) *model.ProximaStreamObject {
	if reader.preprocessFunc == nil {
		return &model.ProximaStreamObject{
			Transition: transition,
			Preprocess: nil,
		}
	}
	return &model.ProximaStreamObject{
		Transition: transition,
		Preprocess: model.NewTransitionPreprocessingResult(transition, reader.preprocessFunc),
	}
}

func (reader *StreamReader) GetStreamChan() (<-chan *model.ProximaStreamObject, <-chan error) {
	return reader.output, reader.outputErr
}

func (reader *StreamReader) Start(ctx context.Context, option *StreamConnectionOption) error {
	if !reader.IsConnected() {
		err := reader.Connect()
		if err != nil {
			return err
		}
	}
	if option == nil {
		option = &StreamConnectionOption{}
	}
	if option.IsEmpty() {
		option = NewDefaultStreamConnectionOption()
	}
	if ctx == nil {
		ctx = context.Background()
	}

	var stream <-chan *model.Transition
	var errc <-chan error
	var err error

	if option.IsRpcBased() {
		stream, errc, err = reader.client.GetStreamBasedOnRpc(ctx, model.StreamState{
			StreamID: reader.config.GetStreamID(),
			State:    reader.lastState,
		})
	}
	if option.IsStreamBased() {
		stream, errc, err = reader.client.GetStream(ctx, model.StreamState{
			StreamID: reader.config.GetStreamID(),
			State:    reader.lastState,
		})
	}
	if err != nil {
		return err
	}
	reader.input = &stream
	reader.inputErr = &errc

	if reader.output == nil {
		reader.output = make(chan *model.ProximaStreamObject, reader.config.GetChannelSize())
		reader.outputErr = make(chan error, 1)
	}
	reader.outputCtx = ctx
	reader.streamConnectionOption = *option
	reader.startProcessing()
	return nil
}

func (reader *StreamReader) Reconnect() error {
	err := reader.Close()
	if err != nil {
		return err
	}
	err = reader.Connect()
	if err != nil {
		return err
	}

	return nil
}

func (reader *StreamReader) Connect() error {
	reader.mutex.Lock()
	if reader.IsConnected() {
		reader.mutex.Unlock()
		return nil
	}
	err := reader.client.Connect()
	if err == nil {
		reader.setIsConnected(true)
	}
	reader.mutex.Unlock()
	return err
}

func (reader *StreamReader) Close() error {
	reader.mutex.Lock()
	if !reader.IsConnected() {
		reader.mutex.Unlock()
		return nil
	}
	err := reader.client.Close()
	if err == nil {
		reader.setIsConnected(false)
	}
	reader.mutex.Unlock()
	return err
}

func (reader *StreamReader) restartStream() {
	if reader.input == nil {
		return
	}
	err := reader.Start(reader.outputCtx, &reader.streamConnectionOption)
	if err != nil {
		reader.outputErr <- err
		return
	}
}

func (reader *StreamReader) reconnectAndRestartStream() {
	err := reader.Reconnect()
	if err != nil {
		reader.outputErr <- err
		return
	}
	reader.restartStream()
}

func (reader *StreamReader) startWatchDog() {
	if reader.streamConnectionOption.WatchDogTimeout == 0 {
		return
	}
	if reader.watchdogCancel != nil {
		reader.watchdogCancel()
	}
	var ctx context.Context
	ctx, reader.watchdogCancel = context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				lastCount := reader.processedCount
				time.Sleep(reader.streamConnectionOption.WatchDogTimeout)
				if !reader.IsConnected() {
					continue
				}
				if reader.processedCount == lastCount {
					log.Println("StreamReader watchdog: no new data, reconnecting...")
					reader.Restart()
				}
			}
		}
	}()
}

func (reader *StreamReader) startProcessing() {
	reader.startWatchDog()
	tick := time.NewTicker(reader.streamConnectionOption.ReconnectTime)
	go func() {
		for {
			select {
			case <-reader.outputCtx.Done():
				return
			case err := <-*reader.inputErr:
				reader.outputErr <- err
				return
			case <-tick.C:
				reader.reconnectAndRestartStream()
				return
			case <-reader.restartSync:
				reader.reconnectAndRestartStream()
				return
			case <-reader.stopSync:
				return
			case obj, ok := <-*reader.input:
				if !ok {
					log.Println("Error reading next object:" + reader.config.GetStreamID())
					reader.restartStream()
					return
				}
				reader.pushObject(reader.streamObjForTransition(obj))
			}
		}

	}()
}

func (reader *StreamReader) Restart() {
	reader.restartSync <- struct{}{}
}

func (reader *StreamReader) Stop() {
	reader.stopSync <- struct{}{}
}

func (reader *StreamReader) ReadNext() (model.StreamObject, error) {
	select {
	case <-reader.outputCtx.Done():
		return nil, reader.outputCtx.Err()

	case obj := <-reader.output:
		if obj.Preprocess.Err != nil {
			return nil, obj.Preprocess.Err
		}
		return obj, nil

	case err := <-reader.outputErr:
		return nil, err
	}
}
