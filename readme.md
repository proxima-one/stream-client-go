# Golang Proxima.one StreamDB Client

This library is a client for Proxima Stream Registry and Proxima StreamDB.

## Stream Registry Client
Implements all methods of the Proxima Streams API that is also available at https://streams.api.proxima.one.

```go
streamRegistryClient := stream_registy.NewStreamRegistryClient(stream_registy.Options{
    Endpoint:        "https://streams.api.proxima.one",
    RetryPolicy:     connection.DefaultPolicy(),
    DebugHttpOutput: false,
})
```

## Proxima Stream Client

```go
client := proxima_stream_client.NewProximaStreamClient(
	proxima_stream_client.Options{Registry: registry}, 
)
```
You can use either `streamRegistryClient` from previous example or create `SingleStreamDbRegistry` as a `registry`:
```go
singleRegistryClient := stream_registy.NewSingleStreamDbRegistry("streams.buh.apps.proxima.one:443")
```
SingleRegistryClient is a simple implementation of the `StreamRegistry` interface that always returns the same stream db address.
It can be useful for development purposes but in production you should use `StreamRegistryClient` that will fetch the StreamDB address from the registry.

As you have created the client you can use it to consume a stream. There are some different methods to do so:

### Streaming events
The second method is more suitable for long-running processes that need to consume a stream in a loop.
```go
stream := client.StreamEvents(
    ctx,                           // stream context. When it is cancelled the stream will be closed
    "proxima.eth-main.blocks.1_0", // the name of the stream
    model.ZeroOffset(),
    1000,                          // stream buffer size. Consider increasing it if you have unstable network connection
)
```
Now `stream` is a Go channel with `model.StreamEvent` structs. You can use it in a loop:
```go
for ctx.Err() == nil {
    select {
    case event := <-stream:
        // process event
    case <-ctx.Done():
        return
    }
}
```
Note that the `StreamEvents` will never throw any error. If there is a problem with the connection it will try to reconnect and continue streaming.

### BufferedStreamReader
In some cases you may want to read events in batches in a long-running process. In this case you can use `BufferedStreamReader`:
```go
reader := proxima_stream_client.NewBufferedStreamReader(stream)
for i := 0; ; i++ {
    events := reader.TryRead(50)
    // process event
}
```
It's single `TryRead` method will read at least one event but no more than the specified number of events and return them as a slice.

<b>If there are no events in the stream it will wait until there is at least one available.</b> If stream has more than one event, it will never wait for more events.

### Fetching a number of events
It is useful when you want to fetch a number of events from the stream, but you shouldn't use it for long-running processes.
```go
events, err := client.FetchEvents(
    "proxima.eth-main.blocks.1_0",       // the name of the stream
    model.ZeroOffset(),
    10,                                  // the MAX number of events to fetch
    proxima_stream_client.DirectionNext, // direction can be either Next or Last which means forward or backward
)
```
You can now process `events` just like any other slice of `model.StreamEvent` structs.

Note that the `FetchEvents` method can return a non-nil error.

