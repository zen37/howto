package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

//import "github.com/kr/pretty"

// ValidateInput validates the input
func ValidateInput(i int, p int) error {
	//TO DO  return all errors in one go rather than exit after each one

	//	fmt.Println("validate input ... "+"iterations:", i, "percentage:", p)

	if i == 1 && p > 0 {
		s := "only one iteration and percentage is higher than zero: "
		return errors.New(s + strconv.Itoa(p))
	}

	if i < 0 {
		s := "number of iterations is negative: "
		return errors.New(s + strconv.Itoa(i))
	}
	if p < 0 {
		s := "error percentage is negative: "
		return errors.New(s + strconv.Itoa(p))
	}
	if p > 100 {
		s := "percentage higher than 100: "
		return errors.New(s + strconv.Itoa(p))
	}

	//	fmt.Println("input is valid")
	return nil
}

func PrintList(s []int) {

	fmt.Println("print list ...")

	var str string
	for _, i := range s {
		str += fmt.Sprintf("%v\n", msg(i))
	}
	fmt.Print(str)

}

func msg(i int) string {
	switch i {
	case 0:
		return "SUCCESS"
	case 1:
		return "ERROR"
	default:
		return "UKNOWN"
	}
}

func PrintListAnyArg(t interface{}) {
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)

		for i := 0; i < s.Len(); i++ {
			fmt.Println(s.Index(i))
		}
	}
}
