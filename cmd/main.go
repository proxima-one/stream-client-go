package main

import (
	"context"
	"fmt"
	"github.com/proxima-one/streamdb-client-go/client"
	"github.com/proxima-one/streamdb-client-go/config"
	"github.com/proxima-one/streamdb-client-go/model"
	"time"
)

func readBatches(reader *client.StreamReader) {
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

func main() {
	cfg := config.NewConfigFromFileOverwriteOptions(
		"config.yaml",
		config.WithChannelSize(10000),
		config.WithState(model.Genesis()),
	)

	reader, err := client.NewStreamReader(cfg, client.JsonParsingPreprocessFunc)

	if err != nil {
		fmt.Println(err)
	}

	//_, _, err = reader.StartGrpcStreamChannel(context.Background())
	_, _, err = reader.StartGrpcRpcChannel(context.Background(), 5000)

	// its not safe to use fix count because of server grpc limit
	// todo: check server limit grpc.MaxRecvMsgSize(????)

	if err != nil {
		fmt.Println(err)
	}
	defer reader.Close()

	start := time.Now()
	processingTime := time.Duration(time.Second / 100)

	//read and process
	go func() {
		for {
			data, err := reader.ReadNext()
			//available time for processing without loading buffer is usually from time.Nanosecond * 100 to time.Nanosecond * 1000
			//you can use it to adjust processing time
			time.Sleep(processingTime) //simulate processing load

			if err != nil {
				fmt.Println(err)
				if err.Error() == "context deadline exceeded" {
					fmt.Println("Finish stream")
					break
				}
				panic(err)
			}
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
	go func() {
		for {
			time.Sleep(time.Second)
			count, bufferLoad := reader.Metrics()
			fmt.Printf("Total messages: %d\n", count)
			fmt.Printf("Msg per sec : %f\n", float64(count)/time.Since(start).Seconds())
			fmt.Printf("Buffer load percent: %d\n", bufferLoad)
			fmt.Printf("Available time for processing: %d Nanosecond\n", processingTime)
			fmt.Printf("-----------------------------------------------------\n")
		}
	}()

	for {
		time.Sleep(time.Second * 30)
		fmt.Println("")
		fmt.Println("...running...")
		fmt.Println("")
	}
}
