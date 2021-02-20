package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	s int = 0 //string = "S"
	e int = 1 //string = "E"
)

//Process process the request
func Process(iter int, perc int) ([]int, error) {

	if iter == 0 {
		fmt.Println("0 iterations, no message to print")
		return nil, nil
	}

	err := ValidateInput(iter, perc)
	if err != nil {
		return nil, err
	}

	randomize, l := createList(iter, perc)

	if randomize {
		return randomizeList(l), nil
	}

	return l, nil

}

//create the list of messages
func createList(iter int, perc int) (bool, []int) {

	//TO DO - what's the best way to deal with non round numbers for the counters

	//	fmt.Println("create list ...")

	countE := iter * perc / 100
	countS := iter - countE

	//	if countE != 0 { //no list to be made if counter = 0
	listE := make([]int, countE)
	for i := 0; i < countE; i++ {
		listE[i] = e
	}
	//	}

	//	if countS != 0 { //no list to be made if counter = 0
	listS := make([]int, countS)
	for i := 0; i < countS; i++ {
		listS[i] = s
	}
	//	}

	if countS == 0 {
		return false, listE
	}

	if countE == 0 {
		return false, listS
	}

	//TO  DO performance wise what's better
	// https://stackoverflow.com/questions/37884361/concat-multiple-slices-in-golang
	return true, append(listS, listE...)

}

func randomizeList(s []int) []int {

	//	fmt.Println("randomize list ...")

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	return s
}
