package main

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/proxima-one/streamdb-client-go/v2/pkg/connection"
	"github.com/proxima-one/streamdb-client-go/v2/pkg/proximaclient"
	"time"
)

var streamRegistryClient *proximaclient.StreamRegistryClient

func getStreamEndpoints() {
	offset, err := proximaclient.NewOffsetFromString("15860589-0xc4db4f4a6c48ffb0d5441cb079cfecf50c528ea3190793be04811c6e2076e27b-1667129423000")
	if err != nil {
		panic(err.Error())
	}
	endpoints, err := streamRegistryClient.GetStreamEndpoints("proxima.eth-main.blocks.1_0", offset)
	if err != nil {
		panic(err.Error())
	}
	spew.Dump(endpoints)
}

func findStream() {
	stream, err := streamRegistryClient.FindStream("proxima.eth-main.blocks.1_0")
	if err != nil {
		panic(err.Error())
	}
	spew.Dump(stream)
}

func getStreams() {
	streams, err := streamRegistryClient.GetStreams()
	if err != nil {
		panic(err.Error())
	}
	spew.Dump(streams)
}

func findStreams() {
	streams, err := streamRegistryClient.FindStreams(&proximaclient.StreamFilter{Labels: map[string]string{
		"encoding": "json",
	}})
	if err != nil {
		panic(err.Error())
	}
	spew.Dump(streams)
}

func findOffset() {
	t := time.Unix(1669046625, 0)
	h := int64(1)
	offset, err := streamRegistryClient.FindOffset("proxima.eth-main.blocks.1_0", &h, &t)
	if err != nil {
		panic(err.Error())
	}
	spew.Dump(offset)
}

func testStreamRegistryClient() {
	streamRegistryClient = proximaclient.NewStreamRegistryClient(proximaclient.StreamRegistryClientOptions{
		Endpoint:        "https://streams.api.proxima.one",
		RetryPolicy:     connection.DefaultPolicy(),
		DebugHttpOutput: false,
	})

	println("\n==================================== getStreamEndpoints() ====================================\n")
	getStreamEndpoints()
	println("\n==================================== findStream() ====================================\n")
	findStream()
	println("\n==================================== getStreams() ====================================\n")
	getStreams()
	println("\n==================================== findStreams() ====================================\n")
	findStreams()
	println("\n==================================== findOffset() ====================================\n")
	findOffset()
	println("\n========================================================================\n")
}

func testStreamDbClientFetch() {
	registry := proximaclient.NewSingleStreamDbRegistry("streams.buh.apps.proxima.one:443")
	client := proximaclient.NewStreamClient(proximaclient.Options{Registry: registry})
	events, err := client.FetchEvents(
		"proxima.eth-main.blocks.1_0",
		proximaclient.ZeroOffset(),
		10,
		proximaclient.DirectionNext,
	)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(len(events))
	spew.Dump(events[0])
}

func testStreamDbClientStream() {
	registry := proximaclient.NewSingleStreamDbRegistry("streams.buh.apps.proxima.one:443")
	client := proximaclient.NewStreamClient(proximaclient.Options{Registry: registry})
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	stream := client.StreamEvents(
		ctx,
		"proxima.eth-main.blocks.1_0",
		proximaclient.NewOffset("0x6df54c6aea7df8327b7dfc74eb8615f3f9b8038b51e435d8e42063382ad555bf", 1000, proximaclient.NewTimestamp(1438272137000, nil)),
		1,
	)
	for i := 0; ; i++ {
		event := <-stream
		println(i, event.Offset.String(), event.Timestamp.Time().String())
		time.Sleep(100 * time.Millisecond)
	}
}

func testStreamDbClientStreamWithBufferedReader() {
	registry := proximaclient.NewSingleStreamDbRegistry("streams.buh.apps.proxima.one:443")
	client := proximaclient.NewStreamClient(proximaclient.Options{Registry: registry})
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	stream := client.StreamEvents(
		ctx,
		"proxima.eth-main.blocks.1_0",
		proximaclient.NewOffset("0x6df54c6aea7df8327b7dfc74eb8615f3f9b8038b51e435d8e42063382ad555bf", 1000, proximaclient.NewTimestamp(1438272137000, nil)),
		1000,
	)
	reader := proximaclient.NewBufferedStreamReader(stream)
	for i := 0; ; i++ {
		events := reader.TryRead(50)
		println(i, len(events))
	}
}

func main() {
	// testStreamRegistryClient()
	// testStreamDbClientFetch()
	// testStreamDbClientStream()
	testStreamDbClientStreamWithBufferedReader()
}
