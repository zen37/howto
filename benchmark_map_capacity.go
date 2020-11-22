// using an accurate capacity is faster, without capacity the map must resize on its own
// with capacity: 55.518µs, 78.766µs, 77.45µs, 67.979µs, 74.115µs
// without: 186.35µs, 193.229µs, 177.525µs, 245.313µs, 170.709µs

package main

import (
	"fmt"
	"time"
)

func main() {

	t0 := time.Now()

	values := make(map[int]int, 1000)
	//values := map[int]int{}

	for x := 0; x < 1000; x++ {
		values[x] = x
	}
	//	}

	t1 := time.Now()
	fmt.Println(t1.Sub(t0))
}
