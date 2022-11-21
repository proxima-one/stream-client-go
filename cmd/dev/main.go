package main

import (
	"fmt"
	"github.com/proxima-one/streamdb-client-go/pkg/connection"
	"github.com/proxima-one/streamdb-client-go/pkg/model"
	"github.com/proxima-one/streamdb-client-go/pkg/stream_registy"
	"time"
)

func main() {
	client := stream_registy.NewStreamRegistryClient(stream_registy.Options{
		Endpoint:        "https://streams.api.proxima.one",
		RetryPolicy:     connection.DefaultPolicy(),
		DebugHttpOutput: false,
	})

	offset, err := model.NewOffsetFromString("15860589-0xc4db4f4a6c48ffb0d5441cb079cfecf50c528ea3190793be04811c6e2076e27b-1667129423000")
	if err != nil {
		panic(err.Error())
	}
	endpoints, err := client.GetStreamEndpoints("proxima.eth-main.blocks.1_0", offset)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%#v\n", endpoints)

	stream, err := client.FindStream("proxima.eth-main.blocks.1_0")
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%#v\n", stream)

	streams, err := client.GetStreams()
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%#v\n", streams)

	streams, err = client.FindStreams(&stream_registy.StreamFilter{Labels: map[string]string{
		"encoding": "json",
	}})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%#v\n", streams)

	t := time.Unix(1, 0)
	h := int64(1)
	offset, err = client.FindOffset("proxima.eth-main.blocks.1_0", &h, &t)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%#v\n", offset)
}
