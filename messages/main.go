package main

import (
	"fmt"
	"log"
	"os"
)

var print bool

func init() {
	if len(os.Args) >= 2 {
		if os.Args[1] == "p" {
			print = true
		}
	}
}

func main() {

	s, err := Process(5, 20) //TO DO provide iterations and percentage as command line arguments
	if err != nil {
		log.Fatal(err)
	}
	if print {
		PrintList(s)
	} else {
		fmt.Println("no print argument provided")
	}
}
