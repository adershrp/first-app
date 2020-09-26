package main

import "fmt"

// GO Lang all PASS BY VALUE.
/**
Usage of Pointers:
1. Instead of sending a large object, just send the address value.
2. Modify the value be de-referencing the address.
*/
func main() {
    a := 100
    fmt.Printf("From main() Value :%v, Type :%T, and Address %v\n", a, a, &a)
    fmt.Println()
    foo(a)
    fmt.Printf("From main() after foo() Value :%v, Type :%T, and Address %v\n", a, a, &a)
    fmt.Println()
    bar(&a)
    fmt.Printf("From main() after bar() Value :%v, Type :%T, and Address %v\n", a, a, &a)
}

func foo(x int) {
    fmt.Printf("From foo(x) Value :%v, Type :%T, and Address %v\n", x, x, &x)
    x = 89
    fmt.Printf("From foo(x) after changing Value :%v, Type :%T, and Address %v\n", x, x, &x)
    fmt.Println()
}

func bar(x *int) {
    fmt.Printf("From bar(x) Value :%v, Type :%T, and Address %v\n", *x, x, x)
    *x = 89
    fmt.Printf("From bar(x) after changing Value :%v, Type :%T, and Address %v\n", *x, x, x)
    fmt.Println()
}
