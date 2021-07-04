package main

import "fmt"

func main() {

	//subslice1()
	//subslice2()
	subslice3()
	fmt.Println("use full slice expression")
	subslice4()

}

//never append with a sublice or use a full slice expression
func subslice3() {
	x := []int{11, 22, 3, 44}
	y := x[:2]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("capacity x is", cap(x), " capacity y is", cap(y))

	y = append(y, 30)
	//y has the length 2, but the capacity is set to 4, the same as x.
	//Since the capacity is 4, appending onto the end of y puts the value in the third position of x.
	fmt.Println("x:", x)
	fmt.Println("y:", y)
}

func subslice4() {
	x := []int{11, 22, 3, 44, 55}
	//use full slice expression, last position :2 for example indicates the last position
	//in the parent's slice capacity that is available for the subslice
	y := x[:2:2]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("capacity x is", cap(x), " capacity y is", cap(y))

	y = append(y, 30)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
}

//take a slice from a slice does not make a copy of the data,
// instead two variables are created that share memory
func subslice2() {
	x := []int{1, 2, 3, 4, 5}
	y := x[:2]
	z := x[3:]
	fmt.Println("x is", x)
	fmt.Println("y is  x[:2]", y)
	fmt.Println("z is x[3:]", z)

	fmt.Println("change first element in x to 10")
	x[0] = 10
	print(x, y, z)
	fmt.Println("change last element in z to 55")
	z[1] = 55
	print(x, y, z)

}

// slicing slices = create slice from a slice
func slice1() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("x is ", x)
	fmt.Println("x[:2] is ", y)
	fmt.Println("x[1:] is ", z)
	fmt.Println("x[1:3] is ", d)
	fmt.Println("x[:] is", e)
}

func print(x []int, y []int, z []int) {
	fmt.Println("x is", x)
	fmt.Println("y is", y)
	fmt.Println("z is", z)
}
