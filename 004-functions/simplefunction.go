package main

import "fmt"

func main() {
    sayHello("Go")
    sayHello2("Hello", "Method 2")
    greeting := "Hello"
    name := "Go"
    sayHello3(&greeting, &name)
    fmt.Println(name)
    s := sum(1, 2, 3, 4, 5, 6, 7)
    fmt.Println("Sum", s)
    t := sum2(1, 2, 3, 4, 5, 6, 7)
    fmt.Println("Sum", *t)
    u := sum3(3, 5, 6)
    fmt.Println("Named Return", u)
    
    c, err := divide(5, 1)
    if nil != err {
        fmt.Println(err)
        return
    }
    fmt.Println("Division", c)
    
    add, sub, mul, div, err := calculator(10, 5)
    if nil != err {
        fmt.Println("Error", err)
    }
    fmt.Println(add, sub, mul, div)
    
    add, sub, mul, div = calculatorNamed(100, 14)
    fmt.Println(add, sub, mul, div)
}

// multiple return
func calculatorNamed(a, b int) (add int, sub int, mul int, div int) {
    add = a + b
    sub = a - b
    mul = a * b
    div = a / b
    return
}

// multiple return
func calculator(a, b int) (int, int, int, int, error) {
    if 0 == a || 0 == b {
        return 0, 0, 0, 0, fmt.Errorf("Cannot accept Zero")
    }
    return a + b, a - b, a * b, a / b, nil
}

// multiple return
// returning error
func divide(a, b float32) (float32, error) {
    if b == 0 {
        return b, fmt.Errorf("Cannot divide by Zero")
    }
    return a / b, nil
}

// named return variable
// specify name of the return variable, no need to init again or mention the attribute as part of return
func sum3(values ...int) (result int) {
    for _, value := range values {
        result += value
    }
    return
}

// returning pointer instead of actual
func sum2(values ...int) *int {
    fmt.Println("Input values", values)
    result := 0
    for _, value := range values {
        result += value
    }
    return &result
}

// vargs
// returning value
func sum(values ...int) int {
    fmt.Println("Input values", values)
    result := 0
    for _, value := range values {
        result += value
    }
    return result
}

// simple function with one arg
func sayHello(msg string) {
    fmt.Println("Hello", msg, "!")
}

// here if both arg are same data type, only last one need the data type
func sayHello2(msg1, msg2 string) {
    fmt.Println(msg1, msg2)
}

// passing the reference (pointers), instead of the copy of the data
// for map and slice this is not required.
func sayHello3(greeting, name *string) {
    fmt.Println(*greeting, *name)
    *name = "New Name"
    fmt.Println(*name)
}
