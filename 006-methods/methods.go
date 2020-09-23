package main

import (
	"fmt"
)

func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
	fmt.Println(g.name)

	g2 := greeter{
		greeting: "Good Morning",
		name:     "Go",
	}
	g2.greetPointer()
	fmt.Println(g2.name)

	fmt.Println("Return", g2.greetReturn())
}

type greeter struct {
	greeting string
	name     string
}

//function with a context called method
//copy of greeter object
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.name = "GO!"
}

//function with a context called method
//pointer with greeter object
func (g *greeter) greetPointer() {
	fmt.Println(g.greeting, g.name)
	g.name = "GO!"
}

//function with a context called method
func (g *greeter) greetReturn() string {
	return fmt.Sprintln(g.greeting, g.name)
}
