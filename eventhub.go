package main

import (
	"context"
	"fmt"
	"log"
	"time"

	eventhub "github.com/Azure/azure-event-hubs-go"
)

func main() {
	connStr := "........"
	hub, err := eventhub.NewHubFromConnectionString(connStr)

	if err != nil {
		fmt.Println(err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// send a single message into a random partition

	now := time.Now()
	//str := `{"name": now }`
	s := now.Format("2006-01-02 15:04:05 Monday")

	err = hub.Send(ctx, eventhub.NewEventFromString(s))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(s)

	// get info about the hub, particularly number and IDs of partitions
	info1, err := hub.GetRuntimeInformation(ctx)
	if err != nil {
		log.Fatalf("failed to get runtime info: %s\n", err)
	}
	fmt.Printf("Partition count: %v\nPath: %v\nCreated At: %v\nPartition IDs %v\n", info1.PartitionCount, info1.Path, info1.CreatedAt, info1.PartitionIDs)
	//log.Printf("partition IDs: %s\n", info1.PartitionIDs)

	// get info about the hub, particularly number and IDs of partitions
	info2, err := hub.GetPartitionInformation(ctx, "0")
	if err != nil {
		log.Fatalf("failed to get runtime info: %s\n", err)
	}
	fmt.Printf("HubPath : %v\nPartitionID: %v\nBeginningSequenceNumber: %v\nLastSequenceNumber: %v\n", info2.HubPath, info2.PartitionID, info2.BeginningSequenceNumber, info2.LastSequenceNumber)

	/*
		handler := func(c context.Context, event *eventhub.Event) error {
			fmt.Println(string(event.Data))
			return nil
		}


			// listen to each partition of the Event Hub
			runtimeInfo, err := hub.GetRuntimeInformation(ctx)
			if err != nil {
				fmt.Println(err)
				return
			}

			for _, partitionID := range runtimeInfo.PartitionIDs {
				// Start receiving messages
				//
				// Receive blocks while attempting to connect to hub, then runs until listenerHandle.Close() is called
				// <- listenerHandle.Done() signals listener has stopped
				// listenerHandle.Err() provides the last error the receiver encountered
				//listenerHandle, err := hub.Receive(ctx, partitionID, handler, eventhub.ReceiveWithLatestOffset())
				_, err := hub.Receive(ctx, partitionID, handler, eventhub.ReceiveWithLatestOffset())

				if err != nil {
					fmt.Println(err)
					return
				}
			}

			// Wait for a signal to quit:
			signalChan := make(chan os.Signal, 1)
			signal.Notify(signalChan, os.Interrupt, os.Kill)
			<-signalChan

			err = hub.Close(context.Background())
			if err != nil {
				fmt.Println(err)
			}
	*/
}
