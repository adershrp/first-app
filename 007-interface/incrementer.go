package main

import (
	"fmt"
)

func main() {
	myInt := IntCounter(0)
	myFloat := FlotCounter(0)
	var intCounter Counter = &myInt
	var floatCounter Counter = &myFloat
	for i := 0; i < 10; i++ {
		fmt.Println(intCounter.Count())
		fmt.Println(floatCounter.Count())
	}
}

type Counter interface {
	Count() int
}

type IntCounter int
type FlotCounter int

func (ic *IntCounter) Count() int {
	*ic++
	return int(*ic)
}
func (ic *FlotCounter) Count() int {
	*ic--
	return int(*ic)
}
