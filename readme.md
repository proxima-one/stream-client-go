# Golang client for proxima streams:

### 1) Proxima streams client (ProximaClient)

- Provides basic access to grpc endpoint (no additional logic for handling errors and retries or state management)
- GetTransitionsAfter(...) returns requested amount transitions after given state
- GetStreams(...) return channel for transitions and channel for errors based on GRPC stream
- GetStreamBasedOnRpc(...) same as GetStreams(...) but based on RPC calls instead of GRPC stream

```go
	//setup connection
	cfg := config.NewConfigFromFileOverwriteOptions(
		"config.yaml",
		config.WithChannelSize(10000),
		config.WithState(model.Genesis()),
	)
    client := NewProximaClient(cfg)
    if /*request-response model*/() {
             stream, errc, err = reader.client.GetStreamBasedOnRpc(ctx, model.StreamState{
             StreamID: reader.config.GetStreamID(),
             State:    reader.lastState,
            })
        }
	if /*streaming model*/() {
             stream, errc, err = reader.client.GetStream(ctx, model.StreamState{
             StreamID: reader.config.GetStreamID(),
             State:    reader.lastState,
	    })
        }
		
    //read stream
	for {
        select {
        case <-ctx.Done():
            return
        case err := <-errc:
            if err != nil {
                return
            }
        case transition, ok := <-stream:
            //process transition
        }
    }
```

### 2) Proxima streams reader (StreamReader)

- Provide easy access to streams based on ProximaClient
- Handle state management (state is stored in memory)
- can easily apply async function to every transition for example: Json Parsing for payload
- trying to restore connection in case of not getting messages from stream for some time period
```go

//StreamConnectionOption struct {
//Type            string //type of connection (grpc, rpc) StreamConnectionOptionTypeRpc, StreamConnectionOptionTypeStream
//ReconnectTime   time.Duration // restart connection every ReconnectTime
//WatchDogTimeout time.Duration // restart connection if WatchDogTimeout is exceeded
//}

reader := client.NewStreamReader(cfg, client.JsonParsingPreprocessFunc)
//you can write your own map function instead JsonParsingPreprocessFunc  fn(transition *model.Transition) -> (any, error)
startParams := client.NewDefaultStreamConnectionOption()
startErr := reader.Start(context.Background(), startParams)


for {
    data, err := reader.ReadNext()
	// check error 
	// if error is not nil then processing stream is stopped 
	// you need to call Start again if error is not critical
	// continueErr := reader.Start(context.Background(), startParams)
	mapValue, err := data.Preprocess.PreprocessingResult()
	//get access to required data from stream item
	//process data
	//...
}
```