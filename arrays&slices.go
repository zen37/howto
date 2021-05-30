package main

import "fmt"

func main() {

	// the array type is defined by its size + type of elements
	// by default an array is zero valued for ints is  0
	var a1 [3]int
	print("array [3]int", a1)
	var s1 []int
	print("slice []int", s1)

	var a2 = [3]int{1, 2, 3}
	fmt.Println("array [3]int{1, 2, 3} -> ", a2)
	var s2 = []int{1, 2, 3}
	fmt.Println("slice []int{1, 2, 3} -> ", s2)

	//sparse array, most elements are set to their 0 values, can specify only the indices with values
	var a3 = [7]int{1, 4: 7, 8}
	print("array [7]int{1, 4: 7, 8}", a3)
	var s3 = []int{1, 4: 7, 8}
	print("slice []int{1, 4: 7, 8}", s3)

	//compare arrays - only arrays of the same type can be compared
	var x = [...]int{1, 2, 3, 4}
	var y = [4]int{1, 2, 3, 4}
	fmt.Println(x == y)

	//multidimensional arrays
	//var m [2][4]int // this is an array of length 2, whose type is an array of ints of length 4
	//print(m)
}

func print(text string, x interface{}) {
	fmt.Println(text, " -> ", x)
}
