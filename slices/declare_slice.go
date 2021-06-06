package main

import "fmt"

func main() {

	declare_slice()

}

//declare and append to slice

func declare_slice() {
	//create slice
	//using an empty slice literal:
	// 		this creates a zero-length slice, which is non-nil (comparing it to nil returns false).

	var y = []int{}
	print(y)
	fmt.Println(y == nil) // false

	// without declaring a literal
	var x []int
	print(x)
	fmt.Println(x == nil) //true

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

}

func print(x []int) {
	fmt.Println("slice is: ", x, "len =", len(x), "cap =", cap(x))
}
