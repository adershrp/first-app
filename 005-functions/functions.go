package main

import (
    "fmt"
)

// anonymous function
func main() {
    // example1()
    // example2()
    var sum = func(a, b int) int {
        return a + b
    }
    s := sum(1, 3)
    fmt.Println(s)
    
}

// functions as variable
func example2() {
    var divide func(a, b float32) (float32, error)
    divide = func(a, b float32) (float32, error) {
        if b == 0 {
            return b, fmt.Errorf("Cannot Divide by Zero")
        }
        return a / b, nil
    }
    d, err := divide(4, 0)
    if nil != err {
        fmt.Println(err)
        return
    }
    fmt.Println("Division", d)
}

// anonymous function
func example1() {
    // on sync
    for i := 0; i < 5; i++ {
        func() {
            fmt.Println(i)
        }()
    }
    // incase of async
    for i := 10; i < 15; i++ {
        func(i int) {
            fmt.Println(i)
        }(i)
    }
    
    f := func(msg string) {
        fmt.Println("Hello", msg)
    }
    f("Go")
}
