// a _goroutine_ is a lightweight thread of execution.

package main

import (
	"fmt"
	"sync"
	"time"
)

func f1(from string) {

	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		fmt.Println(from, ":", i)
	}
}

func f2(from string, j int, wg *sync.WaitGroup) {
	defer wg.Done() //on return, notify the WaitGroup that we’re done
	for i := 0; i < 3; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println(from, ":", i, " time=", j)
	}

}

func main() {

	var wg sync.WaitGroup //declare the sync.WaitGroup

	// Suppose we have a function call `f(s)`. Here's how
	// we'd call that in the usual way, running it synchronously.
	f1("direct")

	// To invoke this function in a goroutine, use `go f(s)`.
	//This new goroutine will execute concurrently with the calling one.

	for j := 1; j <= 3; j++ {
		wg.Add(1)
		time.Sleep(3 * time.Second)
		go f2("goroutine", j, &wg)
	}

	// You can also start a goroutine for an anonymous function call.
	go func(msg string) {
		wg.Add(1)
		time.Sleep(4 * time.Second)
		fmt.Println(msg)
		wg.Done()
	}("going")

	// Our two function calls are running asynchronously in separate goroutines now.
	// Wait for them to finish (for a more robust approach, use a [WaitGroup](waitgroups)).
	//time.Sleep(1 * time.Second)
	wg.Wait() //tell code to wait on the WaitGroup queue to reach zero before proceeding
	fmt.Println("we are done")
}
