//https://gist.github.com/arxdsilva/7392013cbba7a7090cbcd120b7f5ca31

//https://play.golang.org/p/c7qZDlPw-FB

package main

import (
	"fmt"
)

func main() {
	required := []string{"hyx", "ad", "b", "c", "dwwedew"}
	provided := []string{"c", "d", "e", "ad"}
	diff := compare(required, provided)
	fmt.Println("it is required but not provided:", diff)
}

func compare(a, b []string) []string {
	for i := len(a) - 1; i >= 0; i-- {
		for _, vD := range b {
			if a[i] == vD {
				a = append(a[:i], a[i+1:]...)
				break
			}
		}
	}
	return a
}
