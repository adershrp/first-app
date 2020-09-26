package main

import "fmt"

func main() {
    var a = 100
    fmt.Printf("Type of variable %T\t and Value %v\n", a, a)
    fmt.Printf("Type of pointer %T\t and Vaule %v\n", &a, &a)
    // dereferencig the pointer
    fmt.Printf("Type of variable %T\t and Value %v\n", *&a, *&a)
    
    // adding & ( and address) will provide memory location of the variable
    // adding * gives you the value stored in that address.
    // type of the variable and its type of the address for the same variable is different.
    // For example, from above type of a and type of &a are int and *int respectively.
    
    b := &a // assigning the pointer value
    fmt.Printf("Type of variable %T\t and value %v\n", b, b)
    // de-referencing using * operator
    fmt.Printf("Type of variable %T\t and value %v\n", *b, *b)
    
    // updating the value at the address held by variable b -- this will update variable a also.
    *b = 200
    fmt.Printf("Type of variable %T\t and value %v\n", *b, *b)
    fmt.Printf("Type of variable %T\t and Value %v\n", a, a)
    
    x := 19
    y := x
    fmt.Printf("Value of X :%v and Value of Y :%v\n", x, y)
    fmt.Printf("Address of X :%v and Address of Y :%v\n", &x, &y)
    x = 18
    y = 29
    // here values of both X and Y change independently
    fmt.Printf("Value of X :%v and Value of Y :%v\n", x, y)
    fmt.Printf("Address of X :%v and Address of Y :%v\n", &x, &y)
    
}
