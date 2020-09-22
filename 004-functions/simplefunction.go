package main

import "fmt"

func main() {
	sayHello("Go")
	sayHello2("Hello", "Method 2")
	greeting := "Hello"
	name := "Go"
	sayHello3(&greeting, &name)
	fmt.Println(name)
	sum(1, 2, 3, 4, 5, 6, 7)
}

//vargs
func sum(values ...int) {
	fmt.Println("Input values", values)
	result := 0
	for _, value := range values {
		result += value
	}
	fmt.Println("Sum of values", result)
}

// simple function with one arg
func sayHello(msg string) {
	fmt.Println("Hello", msg, "!")
}

//here if both arg are same data type, only last one need the data type
func sayHello2(msg1, msg2 string) {
	fmt.Println(msg1, msg2)
}

//passing the reference (pointers), instead of the copy of the data
//for map and slice this is not required.
func sayHello3(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "New Name"
	fmt.Println(*name)
}
