package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	eventhub "github.com/Azure/azure-event-hubs-go"
)

const (
	connEventHub1 string =  //
	connEventHub2 string = //

func main() {
	connStr := connEventHub2 // connEventHub1
	hub, err := eventhub.NewHubFromConnectionString(connStr)

	if err != nil {
		fmt.Println(err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// get info about the hub, particularly number and IDs of partitions
	info1, err := hub.GetRuntimeInformation(ctx)
	if err != nil {
		log.Fatalf("failed to get runtime info: %s\n", err)
	}
	fmt.Printf("Partition count: %v\nPath: %v\nCreated At: %v\nPartitionIDs %v\n", info1.PartitionCount, info1.Path, info1.CreatedAt, info1.PartitionIDs)

	fmt.Println()

	// get info about the hub, particularly number and IDs of partitions
	info2, err := hub.GetPartitionInformation(ctx, "0")
	if err != nil {
		log.Fatalf("failed to get runtime info: %s\n", err)
	}
	fmt.Printf("HubPath : %v\nPartitionID: %v\nBeginningSequenceNumber: %v\nLastSequenceNumber: %v\nLastEnqueuedOffset: %v\nLastEnqueuedTimeUtc: %v\n",
		info2.HubPath, info2.PartitionID, info2.BeginningSequenceNumber, info2.LastSequenceNumber, info2.LastEnqueuedOffset, info2.LastEnqueuedTimeUtc)

	fmt.Println()

	// send a single message into a random partition

	//str := `{"name": now }`
	//s := now.Format("2006-01-02 15:04:05 Monday")

	for i := 0; i <= 5; i++ {
		var s string

		now := time.Now().String()
		if i != 5 {
			s = "this is bit IT" + " " + strconv.Itoa(i) + " " + now //Format("2006-01-02 15:04:05 Monday")
		} else {
			s = "this is DONE signal " + now //.Format("2006-01-02 15:04:05 Monday")
		}
		err = hub.Send(ctx, eventhub.NewEventFromString(s))
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(5 * time.Second)
		fmt.Println(s)
	}
	//**** read from evenhub ****
	/*
		fmt.Println("read from EventHub")

		exit := make(chan struct{})
		handler := func(ctx context.Context, event *eventhub.Event) error {
			text := string(event.Data)
			fmt.Println(text)
			exit <- struct{}{}
			return nil
		}

		for _, partitionID := range info1.PartitionIDs {
			_, err = hub.Receive(ctx, partitionID, handler)
		}

		// wait for the first handler to get called with "Hello World!"
		<-exit
		err = hub.Close(ctx)
		if err != nil {
			fmt.Println(err)
			return
		}
	*/
}
