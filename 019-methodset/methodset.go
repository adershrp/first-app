package main

import (
    "fmt"
    "math"
)

func main() {
    // c1 is regular type - no pointer
    // c1 non pointer values work with non pointer arguments
    c1 := circle{radius: 10}
    // calcArea(c1)
    
    // c2 pointer values work with non pointer arguments
    calcArea(&c1)
    calcArea2(&c1)
}

func calcArea2(s shape) {
    fmt.Printf("Type %T\t\n", s)
    fmt.Println("Calculated Area :", s.area2())
}

func calcArea(s shape) {
    fmt.Printf("Type %T\t\n", s)
    fmt.Println("Calculated Area :", s.area())
}

type shape interface {
    area() float32
    area2() float32
}

type circle struct {
    radius float32
}

// Receiver without Pointer
// This can be called via passing pointer and non pointer values
func (c circle) area() float32 {
    return math.Pi * c.radius * c.radius
}

// Receiver with pointer
// This can be called via passing only pointer values
func (c *circle) area2() float32 {
    return math.Pi * c.radius * c.radius
}
