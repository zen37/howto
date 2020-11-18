package main

type BigInterface interface {
	A()
	B()
	C()
}

type one struct{}

type two struct {
	i int
	s string
}

func (one) A() {

}

func (one) B() {

}

func (two) A() {

}
func (two) B() {

}

func Test(bi BigInterface) {

}

func main() {

	var x one
	Test(x)

	var y two
	Test(y)

}
