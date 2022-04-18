package main

import (
	"context"
	"fmt"
	"github.com/proxima-one/pocs/stream-db-endpoint/client"
	"github.com/proxima-one/pocs/stream-db-endpoint/config"
	"github.com/proxima-one/pocs/stream-db-endpoint/model"
	"time"
)

func readBatch(reader *client.StreamReader) {
	start := time.Now()
	processed := 0
	for {
		ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
		transitions, _ := reader.FetchNextTransitions(ctx, 1000)
		if len(transitions) == 0 {
			fmt.Printf("total messages %d\n", processed)
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
	}
	for {
		transition := <-data
		processed++
		if transition.IsEmpty() {
			fmt.Printf("total messages %d\n", processed)
			fmt.Printf("Finish stream")
			break
		}
		if processed%10000 == 0 {
			fmt.Println(transition.Event.Timestamp)
			elapsed := time.Since(start)
			fmt.Printf("Batched Stream processed %f transitions per sec", float64(processed)/elapsed.Seconds())
		}
	}
}

func readStream(reader *client.StreamReader) {
	start := time.Now()
	processed := 0
	data, _, err := reader.GetRawStreamFromState(context.Background(), model.Genesis(), 10000)
	if err != nil {
		fmt.Println(err)
	}
	for {
		transition := <-data
		processed++
		if transition.IsEmpty() {
			fmt.Printf("Total messages %d\n", processed)
			fmt.Printf("Finish stream")
			break
		}
		if processed%10000 == 0 {
			fmt.Println(transition.Event.Timestamp)
			elapsed := time.Since(start)
			fmt.Printf("Raw Stream processed %f transitions per sec", float64(processed)/elapsed.Seconds())
		}
	}
}

func main() {
	config := config.NewConfigFromYamlFile("config.yaml")

	go func() {
		reader := client.NewStreamReader(*config, model.Genesis())
		reader.Connect()
		defer reader.Disconnect()
		readBatch(reader)
	}()

	go func() {
		reader := client.NewStreamReader(*config, model.Genesis())
		reader.Connect()
		defer reader.Disconnect()
		readStream(reader)
	}()

	go func() {
		reader := client.NewStreamReader(*config, model.Genesis())
		reader.Connect()
		defer reader.Disconnect()
		readBatchedStream(reader)
	}()
	for {
		time.Sleep(time.Second * 10)
		fmt.Println("")
		fmt.Println("...running...")
		fmt.Println("")
	}
}
