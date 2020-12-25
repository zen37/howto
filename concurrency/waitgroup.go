/*
Calhoun.io
A WaitGroup waits for a collection of goroutines to finish.
The main goroutine calls Add to set the number of goroutines to wait for.
Then each of the goroutines runs and calls Done when finished.
At the same time, Wait can be used to block until all goroutines have finished.

We can do everything a sync.WaitGroup does with a channel, so why the new type?
The short answer is that sync.WaitGroup is a bit clearer
when we don't care about the data being returned from the goroutine.
It signifies to other developers that we just want to wait for a group of goroutines to complete,
whereas a channel communicates that we are interested in results from those goroutines.

https://play.golang.org/p/9wdS5HaghRx
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func notify(services ...string) {

	var wg sync.WaitGroup //declare the sync.WaitGroup

	for _, service := range services {
		wg.Add(1) //add to the WaitGroup queue
		go process(service, &wg)
	}

	wg.Wait() //tell code to wait on the WaitGroup queue to reach zero before proceeding
	fmt.Println("All services notified!")
}

func main() {
	notify("Service 1", "Service 2", "Service 3")
}

func process(s string, wg *sync.WaitGroup) {
	fmt.Printf("Starting to notifing %s...\n", s)
	time.Sleep(3 * time.Second)
	fmt.Printf("Finished notifying %s...\n", s)
	wg.Done() //mark items in the queue as done
}
