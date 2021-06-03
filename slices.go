package main

import "fmt"

func main() {

	//declare_slice()

	make_slice()

}

//declare and append to slice

func declare_slice() {
	//create slice
	//using an empty slice literal:
	// 		this creates a zero-length slice, which is non-nil (comparing it to nil returns false).
	/*
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
	*/
}

/*
https://learning.oreilly.com/library/view/learning-go/9781492077206/ch03.html
If you have a good idea of how large your slice needs to be,
but don’t know what those values will be when you are writing the program, use make.
*/
// make, allows to specify the type, length, and, optionally, the capacity.
func make_slice() {

	fmt.Println("Use a non-zero length or zero length and capacity?")
	fmt.Println(" - If you are using a slice as a buffer, then specify a nonzero length.")

	opt2 := ` - If you are sure you know the exact size you want,
	you can specify the length and index into the slice to set the values.
	This is often done when transforming values in one slice and storing them in a second.
	The downside to this approach is that if you have the size wrong,
	you’ll end up with either zero values at the end of the slice or
	a panic from trying to access elements that don’t exist.`

	fmt.Println(opt2)

	opt3 := ` - In other situations, use make with a zero length and a specified capacity. 
	This allows you to use append to add items to the slice. 
	If the number of items turns out to be smaller, you won’t have an extraneous zero value at the end. 
	If the number of items is larger, your code will not panic.`

	fmt.Println(opt3)

	s := `Conclusion: The Go community is split between the second and third approaches. 
	I [Jon Bodner] personally prefer using append with a slice initialized to a zero length. 
	It might be slower in some situations, but it is less likely to introduce a bug.`

	fmt.Println(s)
	fmt.Println()

	//make an int slice of 3 elements that are 0, len is 3 and capacity is 3
	x := make([]int, 3)
	print(x, "make([]int, 3)")

	// 10 is appended to the end of the slice, len and cap are increased
	x = append(x, 10)
	print(x, "append(x, 10)")

	// specify initial capacity
	x = make([]int, 3, 32)
	print(x, "make([]int, 3, 32)")

	// length 0
	x = make([]int, 0, 10)
	print(x, "make([]int, 0, 10)")

	x = append(x, 1, 2, 3)
	print(x, "append(x, 1, 2, 3)")

}

func print(x []int, what string) {
	fmt.Println(what+" - "+"slice is: ", x, "len =", len(x), "cap =", cap(x))
}
