package main

import (
    "fmt"
)

func main() {
    // * defining call back func as anonymous
    cb1 := func(x int) string {
        return fmt.Sprint("Result :\t", x)
    }
    
    // * defining call back func as anonymous
    cb2 := func(x int) string {
        return fmt.Sprint("Answer :\t", x)
    }
    
    // same add has two callback
    fmt.Println(add(10, 20, cb1))
    fmt.Println(add(10, 20, cb2))
    
    fmt.Println(sub(10, 20, cb2))
    fmt.Println(mul(10, 20, cb1))
    fmt.Println(div(10, 20, cb2))
    
}

// add with callback
func add(x, y int, c func(x int) string) string {
    return c(x + y)
}

// sub with callback
func sub(x, y int, c func(x int) string) string {
    return c(x - y)
}

// mul with callback
func mul(x, y int, c func(x int) string) string {
    return c(x * y)
}

// div with callback
func div(x, y int, c func(x int) string) string {
    return c(x / y)
}
