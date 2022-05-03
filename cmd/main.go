package main

import (
	"context"
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
		transition, ok := <-data
		processed++
		if transition == nil || !ok {
			fmt.Printf("total messages %d\n", processed)
			fmt.Printf("Finish stream, breaking processing, or you can wait for the next batch")
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
	ctx, _ := context.WithTimeout(context.Background(), time.Second*100)
	data, _, err := reader.GetRawStream(ctx, 10000)
	if err != nil {
		fmt.Println(err)
	}
	for {
		transition, ok := <-data
		if !ok {
			fmt.Printf("total messages %d\n", processed)
			fmt.Printf("Finish stream")
			break
		}
		processed++
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
