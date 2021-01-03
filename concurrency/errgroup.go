/*

www.calhoun.ioit add 
The Go standard library follows a fairly strict backwards compatibility promise.
As a developer this is a good thing, because it means the packages there won't randomly change and break your code, but it has downsides.
The /x/ stands for external, and the packages inside it are still part of the Go project, but are external to the Go standard library. This allows the developers to follow a looser set of backwards compatbility guidelines on packages where the strict guidelines of the standard library don't make sense.
You can find more of these packages at https://golang.org/pkg/#subrepo​
Aside from needing to `go get` these packages, we can basically treat these packges as if they are part of the standard library

*/

package main

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	notify("Service-1", "Service-2", "Service-3")
}

func notify(services ...string) error {
	var g errgroup.Group

	for _, service := range services {
		// Copy the value from the service variable into a local variable to
		// avoid the bug described at https://www.calhoun.io/gotchas-and-common-mistakes-with-closures-in-go/#variables-declared-in-for-loops-are-passed-by-reference
		s := service
		g.Go(func() error {
			fmt.Printf("Starting to notifing %s...\n", s)
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			fmt.Printf("Finished notifying %s...\n", s)
			return nil // or a real error if we had one
		})
	}

	err := g.Wait()
	if err != nil {
		return fmt.Errorf("notify: %w", err)
	}
	fmt.Println("All services notified successfully!")
	return nil
}
