package main

import "fmt"

func main() {

	//create slice
	//using an empty slice literal:
	// 		this creates a zero-length slice, which is non-nil (comparing it to nil returns false).
	var y = []int{}
	fmt.Println(y == nil) // false
	print(y)

	// without declaring a literal
	var z []int
	fmt.Println(z == nil) //true

	//capacity
	var x []int
	print(x)

	x = append(x, 10)
	print(x)

	x = append(x, 20)
	print(x)

	x = append(x, 30)
	print(x)

	x = append(x, 40)
	print(x)

	x = append(x, 50)
	print(x)

	make_slice()

}

// make, allows to specify the type, length, and, optionally, the capacity.
func make_slice() {

	//make an int slice of 3 elements that are 0, len is 3 and capacity is 3
	x := make([]int, 3)
	print(x)

	// 10 is appended to the end of the slice, len and cap are increased
	x = append(x, 10)
	print(x)

	// specify initial capacity
	x = make([]int, 3, 32)
	print(x)

	// length 0
	x = make([]int, 0, 10)
	print(x)

	x = append(x, 1, 2, 3)
	print(x)

}

func print(x []int) {
	fmt.Println("slice is: ", x, "len =", len(x), "cap =", cap(x))
}
