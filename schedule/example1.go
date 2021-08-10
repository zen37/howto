//https://stackoverflow.com/questions/40364270/running-a-function-periodically-in-go

package main

import (
	"fmt"
	"time"
)

func main() {
	for range time.Tick(time.Second * 60) {
		doSomething()
	}
}

func doSomething() {
	fmt.Println("hello tick", time.Now().Format("2006-01-02 15:04:05"))
}
