package main

import (
	"context"

	"fmt"

	"io/ioutil"

	"log"

	"os"

	"strings"

	"time"

	eventhub "github.com/Azure/azure-event-hubs-go"
)

const (
	connEventHub string = "………connection string …."

	path_order string = "files/pc.json"
)

var eventType string

var additionalInfo string

func init() {

	//fmt.Println(cfg.Conn)

	if len(os.Args) == 2 {

		eventType = os.Args[1]

	}

	if len(os.Args) == 3 {

		eventType = os.Args[1]

		additionalInfo = os.Args[2]

	}

}

func main() {

	hub, err := eventhub.NewHubFromConnectionString(connEventHub)

	if err != nil {

		fmt.Println(err)

		return

	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	// get info about the hub

	info, err := hub.GetRuntimeInformation(ctx)

	if err != nil {

		log.Fatalf("failed to get runtime info: %s\n", err)

	}

	if additionalInfo == "v" {

		fmt.Printf("EntityPath: %v\n", info.Path)

		fmt.Printf("PartitionCount: %v\n", info.PartitionCount)

		fmt.Printf("PartitionIDs: %v\n", info.PartitionIDs)

		fmt.Printf("CreatedAt: %v\n", info.CreatedAt)

	}

	fmt.Printf("EventHub: %v\n", connEventHub[0:53]+"...")

	s, err := sendEvent(ctx, hub)

	if err != nil {

		log.Fatal(err)

	}

	fmt.Println("SUCCESS sending the following data:")

	fmt.Println(s)

}

func sendEvent(ctx context.Context, hub *eventhub.Hub) (string, error) {

	// read file

	var f string

	switch eventType {

	case "o":

		f = path_order

	case "b":

		f = path_ba

	case "t":

		f = path_tax

	case "d":

		f = path_done

	case "r":

		f = path_recs

	default:

		f = path_ba

	}

	data, err := ioutil.ReadFile(f)

	if err != nil {

		return "", err

	}

	s := string(data)

	s = strings.Replace(s, " ", "", -1)

	err = hub.Send(ctx, eventhub.NewEventFromString(s))

	if err != nil {

		return "", err

	}

	//fmt.Println(s)

	return s, nil

}
