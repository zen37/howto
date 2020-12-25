/*
Calhoun.io
https://play.golang.org/p/RpxPoLQRZP-
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	notify("Service 1", "Service 2", "Service 3")
}

func notify(services ...string) {

	ch := make(chan string)

	for _, service := range services {

		t := time.Duration(rand.Intn(8))

		go func(s string) {
			fmt.Printf("Starting to notifing %s...\n", s)
			time.Sleep(t * time.Second)
			ch <- fmt.Sprintf("Finished %s - sleep sec. %d", s, t)
		}(service)
	}

	for i := 0; i < len(services); i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("All services notified!")
}
