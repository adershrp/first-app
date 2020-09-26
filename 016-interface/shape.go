package main

import (
    "fmt"
    "math"
)

func main() {
    // circle
    c := newCircle(6.)
    fmt.Println(info(c))
    
    // square
    s := newSquare(4, 5)
    fmt.Println(info(s))
    
    f := func(nos int) int {
        total := 0
        for i := 0; i < 10; i++ {
            total += i
        }
        return total
    }
    fmt.Println("Sum of 10", f(10))
    sf := somefunc(10)
    fmt.Println(sf())
    
    callback := func(xi []string) string {
        result := "Result"
        for _, v := range xi {
            result += v
        }
        return result
    }
    fmt.Println(someString("hello", callback))
    
}

func someString(s string, callback func([]string) string) string {
    return callback([]string{s, s, s})
}

func somefunc(x int) func() string {
    return func() string {
        return fmt.Sprint(x)
    }
}

func info(s shape) string {
    return fmt.Sprint("Area :", s.area())
}

type shape interface {
    area() float32
}

// type struct
type square struct {
    length, width float32
}

func newSquare(length float32, width float32) *square {
    return &square{length: length, width: width}
}

type circle struct {
    radius float32
}

func newCircle(radius float32) *circle {
    return &circle{radius: radius}
}

func (s square) area() float32 {
    return s.length * s.width
}
func (c circle) area() float32 {
    return math.Pi * c.radius * c.radius
}
