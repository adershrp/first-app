package main

import (
	"fmt"
)

func main() {
	var cal Calculator = Addition{}
	result, err := cal.Calculate(10, 23)
	if nil != err {
		fmt.Println("Error while addition", err)
	}
	fmt.Println("Addition", result)

	cal = Subtraction{}
	result, err = cal.Calculate(10, 23)
	if nil != err {
		fmt.Println("Error while addition", err)
	}
	fmt.Println("Subtraction", result)

	cal = Multiplication{}
	result, err = cal.Calculate(10, 23)
	if nil != err {
		fmt.Println("Error while addition", err)
	}
	fmt.Println("Multiplication", result)

	cal = Division{}
	result, err = cal.Calculate(10, 23)
	if nil != err {
		fmt.Println("Error while addition", err)
	}
	fmt.Println("Division", result)
}

type Calculator interface {
	Calculate(int, int) (int, error)
}

type Addition struct{}
type Subtraction struct{}
type Multiplication struct{}
type Division struct{}

func (add Addition) Calculate(a, b int) (result int, err error) {
	result = a + b
	err = nil
	return
}
func (sub Subtraction) Calculate(a, b int) (result int, err error) {
	result = a - b
	err = nil
	return
}
func (mul Multiplication) Calculate(a, b int) (result int, err error) {
	result = a * b
	err = nil
	return
}
func (div Division) Calculate(a, b int) (result int, err error) {
	result = a / b
	err = nil
	return
}
