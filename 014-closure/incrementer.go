package main

import (
    "fmt"
)

func main() {
	a := incrementer()
	b := incrementer()
	
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func incrementer() func() int {
	// !scope of x is between this block
	var x int
	return func() int {
		x++
		return x
	}
}
