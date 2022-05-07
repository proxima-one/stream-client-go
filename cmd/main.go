package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/proxima-one/streamdb-client-go/client"
	"github.com/proxima-one/streamdb-client-go/config"
	"github.com/proxima-one/streamdb-client-go/model"
	"time"
)

func readBatch(reader *client.StreamReader) {
	start := time.Now()
	processed := 0
	for {
		ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
		transitions, _ := reader.FetchNextTransitions(ctx, 1000)
		if len(transitions) == 0 {
			fmt.Printf("\ntotal messages %d\n", processed)
			fmt.Printf("Finish stream")
			break
		}
		for _, transition := range transitions {
			processed++
			if processed%10000 == 0 {
				fmt.Println(transition.Event.Timestamp)
				elapsed := time.Since(start)
				fmt.Printf("Fetch processed %f transitions per sec", float64(processed)/elapsed.Seconds())
			}
		}
	}
}

func readBatchedStream(reader *client.StreamReader) {
	start := time.Now()
	processed := 0
	data, _, err := reader.GetBatchedStream(context.Background(), 5000, 4000)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		msg, ok := <-data
		transition := msg.Transition
		processed++
		if transition == nil || !ok {
			fmt.Printf("total messages %d\n", processed)
			fmt.Printf("Finish stream, breaking processing, or you can wait for the next batch")
			break
		}
		if processed%40000 == 0 {
			tType, _ := msg.Preprocess.PreprocessingResult()
			fmt.Printf("Last transaction type: %s on state %s\n ", tType, transition.NewState.Id)
			elapsed := time.Since(start)
			fmt.Printf("Batch Stream processed %f transitions per sec \n", float64(processed)/elapsed.Seconds())
			fmt.Printf("Buffer load : %f\n", reader.GetStreamBufferLoad())
		}
	}
}

func readStream(reader *client.StreamReader) {
	data, _, err := reader.GetRawStream(context.Background(), 5000)
	start := time.Now()
	processed := 0

	if err != nil {
		fmt.Println(err)
	}
	for {
		msg, ok := <-data
		transition := msg.Transition
		if !ok {
			fmt.Printf("total messages %d\n", processed)
			fmt.Printf("Finish stream")
			break
		}
		processed++
		if processed%40000 == 0 {
			tType, _ := msg.Preprocess.PreprocessingResult()
			fmt.Printf("Last transaction type: %s on state %s\n ", tType, transition.NewState.Id)
			elapsed := time.Since(start)
			fmt.Printf("Raw Stream processed %f transitions per sec \n", float64(processed)/elapsed.Seconds())
			fmt.Printf("Buffer load : %f\n", reader.GetStreamBufferLoad())
		}
	}
}

func typeFromTransition(transition *model.Transition) (any, error) {
	m := make(map[string]interface{})
	payload := transition.Event.Payload
	err := json.Unmarshal(*payload, &m)
	return m["type"].(string), err
}

func main() {
	config := config.NewConfigFromYamlFile("config.yaml")

	//go func() {
	//	reader, err := client.NewStreamReader(*config, model.Genesis(), nil)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	defer reader.Close()
	//	readBatch(reader)
	//}()

	go func() {
		reader, err := client.NewStreamReader(*config, model.Genesis(), typeFromTransition)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer reader.Close()
		readStream(reader)
	}()

	//go func() {
	//	reader, err := client.NewStreamReader(*config, model.Genesis(), typeFromTransition)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	defer reader.Close()
	//	readBatchedStream(reader)
	//}()
	for {
		time.Sleep(time.Second * 10)
		fmt.Println("")
		fmt.Println("...running...")
		fmt.Println("")
	}
}
