package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	successStr = "SUCCESS"
	errorStr   = "ERROR"
)

func Process(totalMsgs, errPerc int) {

	var currentSuccesses, currentErrors int

	fmt.Println(totalMsgs)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	totalErrs := float64(totalMsgs) * float64(float64(errPerc)/100)

	for _, value := range r.Perm(totalMsgs) {
		if float64(value) < totalErrs {
			currentErrors++
			continue
		}

		currentSuccesses++
	}

	fmt.Printf("SUCCESSES: %d\n", currentSuccesses)
	fmt.Printf("ERRORS: %d\n", currentErrors)
}
