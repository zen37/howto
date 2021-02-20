package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	s int = 0 //TO Do use enums
	e int = 1
)

//Process processes the request
func Process(iter int, perc int) ([]int, error) {

	if iter == 0 {
		fmt.Println("0 iterations, no message to print")
		return nil, nil
	}

	err := ValidateInput(iter, perc)
	if err != nil {
		return nil, err
	}

	randomize, s := createList(iter, perc)

	if randomize {
		return randomizeList(s), nil
	}

	return s, nil

}

//create the list of messages
func createList(iter int, perc int) (bool, []int) {

	//TO DO - what's the best way to deal with non round numbers for the counters

	//	fmt.Println("create list ...")

	list := make([]int, iter)

	countE := iter * perc / 100

	for i := 0; i < countE; i++ {
		list[i] = e
	}

	//fmt.Println(list)

	if countE == 0 || perc == 100 {
		return false, list
	}

	return true, list

}

func randomizeList(s []int) []int {

	//	fmt.Println("randomize list ...")

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
	return s
}
