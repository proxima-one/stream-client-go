package main

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/proxima-one/streamdb-client-go/pkg/connection"
	"github.com/proxima-one/streamdb-client-go/pkg/proxima_stream_client"
	"github.com/proxima-one/streamdb-client-go/pkg/stream_model"
	"github.com/proxima-one/streamdb-client-go/pkg/stream_registy"
	"time"
)

var streamRegistryClient *stream_registy.StreamRegistryClient

func getStreamEndpoints() {
	offset, err := stream_model.NewOffsetFromString("15860589-0xc4db4f4a6c48ffb0d5441cb079cfecf50c528ea3190793be04811c6e2076e27b-1667129423000")
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
	streams, err := streamRegistryClient.FindStreams(&stream_registy.StreamFilter{Labels: map[string]string{
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
	streamRegistryClient = stream_registy.NewStreamRegistryClient(stream_registy.Options{
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
	registry := stream_registy.NewSingleStreamDbRegistry("streams.buh.apps.proxima.one:443")
	client := proxima_stream_client.NewProximaStreamClient(proxima_stream_client.Options{Registry: registry})
	events, err := client.FetchEvents(
		"proxima.eth-main.blocks.1_0",
		stream_model.ZeroOffset(),
		10,
		stream_model.DirectionNext,
	)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(len(events))
	spew.Dump(events[0])
}

func testStreamDbClientStream() {
	registry := stream_registy.NewSingleStreamDbRegistry("streams.buh.apps.proxima.one:443")
	client := proxima_stream_client.NewProximaStreamClient(proxima_stream_client.Options{Registry: registry})
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	stream := client.StreamEvents(
		ctx,
		"proxima.eth-main.blocks.1_0",
		stream_model.NewOffset("0x6df54c6aea7df8327b7dfc74eb8615f3f9b8038b51e435d8e42063382ad555bf", 1000, stream_model.NewTimestamp(1438272137000, nil)),
		1,
	)
	for i := 0; ; i++ {
		event := <-stream
		println(i, event.Offset.String(), event.Timestamp.Time().String())
		time.Sleep(100 * time.Millisecond)
	}
}

func testStreamDbClientStreamWithBufferedReader() {
	registry := stream_registy.NewSingleStreamDbRegistry("streams.buh.apps.proxima.one:443")
	client := proxima_stream_client.NewProximaStreamClient(proxima_stream_client.Options{Registry: registry})
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	stream := client.StreamEvents(
		ctx,
		"proxima.eth-main.blocks.1_0",
		stream_model.NewOffset("0x6df54c6aea7df8327b7dfc74eb8615f3f9b8038b51e435d8e42063382ad555bf", 1000, stream_model.NewTimestamp(1438272137000, nil)),
		1000,
	)
	reader := proxima_stream_client.NewBufferedStreamReader(stream)
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
