package main

import (
	"context"
	"fmt"
	"github.com/proxima-one/streamdb-client-go/client"
	"github.com/proxima-one/streamdb-client-go/config"
	"github.com/proxima-one/streamdb-client-go/model"
)

func main() {
	reader := client.NewFileStreamReader(config.NewConfigWithOptions(
		config.WithState(model.State{Id: ""}), config.WithStreamID("file_stream_reader_example/file.json")),
		func(transition *model.Transition) (any, error) {
			return fmt.Sprintf("Hello, my timestamp is %s", transition.Event.Timestamp), nil
		})
	reader.Start(context.Background(), nil)
	for {
		msg, _ := reader.ReadNext()
		if msg == nil {
			return
		}
		prep, _ := msg.Preprocess.PreprocessingResult()
		fmt.Printf("%+v | %s\n", *msg.Transition, prep.(string))
	}
}
