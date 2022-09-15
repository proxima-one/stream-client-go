package main

import (
	"context"
	"fmt"
	"github.com/proxima-one/streamdb-client-go/client"
	"github.com/proxima-one/streamdb-client-go/config"
	"github.com/proxima-one/streamdb-client-go/model"
	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	"regexp"
	"strconv"
	"time"
)

func getNumberID(id string) string {
	re := regexp.MustCompile("[0-9]+")
	return re.FindString(id)
}

func runLogs(reader *client.StreamReader, processingTime *time.Duration) {
	start := time.Now()
	go func() {
		for {
			time.Sleep(time.Second)
			count, bufferLoad := reader.Metrics()
			fmt.Printf("Total messages: %d\n", count)
			fmt.Printf("Msg per sec : %f\n", float64(count)/time.Since(start).Seconds())
			fmt.Printf("Buffer load percent: %d\n", bufferLoad)
			fmt.Printf("Available time for processing: %d Nanosecond\n", *processingTime)
			fmt.Printf("-----------------------------------------------------\n")
		}
	}()
}

func simulateNetworkProblems(reader *client.StreamReader) {
	go func() {
		for {
			time.Sleep(time.Second * 30)
			reader.Reconnect()
		}
	}()
}

func main() {
	//setup connection
	cfg := config.NewConfigFromFileOverwriteOptions(
		"config.yaml",
		config.WithChannelSize(10000),
		config.WithState(model.Genesis()),
	)

	reader := client.NewStreamReader(cfg, model.JsonParsingPreprocessFunc)
	startParams := client.NewStreamBasedStreamConnectionOption(time.Second * 4)

	startStreamWithRetryFunc := func() {
		for {
			restartErr := reader.Start(context.Background(), startParams)
			if restartErr != nil {
				fmt.Println(restartErr)
				time.Sleep(time.Second)
			} else {
				break
			}
		}
	}
	startStreamWithRetryFunc()

	processingTime := time.Second / 10

	//read and process
	ind := 0

	go func() {
		for {
			data, err := reader.ReadNext()
			if err != nil {
				if status.Code(err) == codes.Unavailable ||
					status.Code(err) == codes.Canceled {
					fmt.Println("Network issues, trying to restart stream")
					startStreamWithRetryFunc()
					continue
				}
				panic(err)
			}

			if data == nil && err == nil {
				panic("There should be data or error")
			}
			newind, _ := strconv.Atoi(getNumberID(data.Transition.NewState.Id))
			if newind-ind > 1 {
				panic("missed item or broken stream")
			}
			ind = newind
			//available time for processing without loading buffer is usually from time.Nanosecond * 100 to time.Nanosecond * 1000
			//you can use it to adjust processing time maybe increase or decrease internal buffer size or save batch size
			time.Sleep(processingTime) //simulate processing load
			if data == nil {
				fmt.Println("Finish stream")
				break
			}
		}
	}()

	//balance processing load //you can use it to adjust processing time for example save batch size
	go func() {
		for {
			time.Sleep(time.Second / 10)
			_, bufferLoad := reader.Metrics()
			if bufferLoad > 95 {
				processingTime = processingTime / 2
			} else {
				adj := 50 - bufferLoad
				processingTime = processingTime + time.Duration(adj)*time.Nanosecond
			}
		}
	}()

	//logs
	runLogs(reader, &processingTime)
	simulateNetworkProblems(reader)

	for {
		time.Sleep(time.Second * 300)
		fmt.Println("")
	}
}
