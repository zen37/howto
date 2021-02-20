package main

import "fmt"

const (
	s string = "SUCCESS"
	e string = "ERROR"
)

func main() {

	alternate(4, 25)

}

func alternate(iter int, perc int) { // use uint?

	//validation check for

	if perc == 0 {
		all_same(s, iter)
		return
	}

	if perc == 100 {
		all_same(e, iter)
		return
	}

	err_max := iter * perc / 100 // what if the result is not a round number

	var err_count int
	for i := 1; i <= iter; i++ {

		err_count = err_count + 1

		if err_count <= err_max {
			fmt.Println(e, i)
		} else {
			fmt.Println(s, i)
		}
		/*
			if i%2 == 0 {
				fmt.Println(s, i)
			} else {
				fmt.Println(e, i)
			}\
		*/
	}
}

func all_same(msg string, iter int) {
	for i := 1; i <= iter; i++ {
		fmt.Println(msg, i)
	}
}
