package main

import "fmt"

func main() {
	sayHello("Go")
	sayHello2("Hello", "Method 2")
	greeting := "Hello"
	name := "Go"
	sayHello3(&greeting, &name)
	fmt.Println(name)
}
func sayHello(msg string) {
	fmt.Println("Hello", msg, "!")
}

func sayHello2(msg1, msg2 string) {
	fmt.Println(msg1, msg2)
}

func sayHello3(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "New Name"
	fmt.Println(*name)
}
