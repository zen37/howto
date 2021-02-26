package main

import (
	"math/rand"
	"sync"
	"time"
)

type counters struct {
	currentSuccesses int
	currentErrors    int
	ctrl             sync.Mutex
}

// const (
//  successStr = "SUCCESS"
//  errorStr   = "ERROR"
// )

func Process(totalMsgs, errPerc int) {
	var c counters
	// fmt.Println(totalMsgs)
	//var currentSuccesses, currentErrors int

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	wg := &sync.WaitGroup{}
	wg.Add(totalMsgs)

	totalErrs := float64(totalMsgs) * float64(float64(errPerc)/100)

	for _, value := range r.Perm(totalMsgs) {
		if float64(value) < totalErrs {
			go func(value int, c *counters) {
				c.ctrl.Lock()
				c.currentErrors++
				c.ctrl.Unlock()
				// fmt.Printf(errorStr)
				// fmt.Printf("error: %d\n", value+1)
				wg.Done()
			}(value, &c)
			//currentErrors++
			continue
		}

		go func(c *counters) {
			c.ctrl.Lock()
			c.currentSuccesses++
			c.ctrl.Unlock()
			//fmt.Println("success")
			wg.Done()
		}(&c)
		//fmt.Printf(successStr)
		//currentSuccesses++
	}
	wg.Wait()

	// fmt.Printf("SUCCESSES: %d\n", currentSuccesses)
	// fmt.Printf("ERRORS: %d\n", currentErrors)
}
