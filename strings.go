package main

import (
	"fmt"
	"strings"
)

func main() {
	// split a string into substrings removing any space characters, including newlines.
	fmt.Println(strings.Fields(" a \t b \n"))

	s := "http://localhost:8080/view/^&"
	res := strings.Split(s, "/")
	fmt.Println(res[len(res)-1])

	fmt.Println(strings.Split(s, "/"))

	fmt.Printf("%q\n", strings.Split(s, "/"))

	fmt.Println(strings.SplitAfter(s, "/"))

	fmt.Println(strings.SplitN(s, "/", 2))
	fmt.Println(strings.SplitN(s, "/", 3))

}
