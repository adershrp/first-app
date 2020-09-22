package main

import "fmt"

func main() {
	example1()
}

func example1() {
	var a int = 185
	// var b *int
	b := &a
	// fmt.Println(a, b) // without pointer prints different memory address
	fmt.Println(a, *b) // with pointer both are referencing same memory location
	a = 215            // can be update using this, both a and b change
	fmt.Println(a, *b)
	*b = 643 // both a and b change
	fmt.Println(a, *b)
}
