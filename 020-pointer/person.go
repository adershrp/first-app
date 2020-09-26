package main

import "fmt"

func main() {
    p1 := person{age: 35, last: "rp", first: "adersh"}
    p1.printNameA()
    p1.printNameB()
    
}

type person struct {
    first string
    last  string
    age   int
}

func (p person) printNameA() {
    fmt.Println(p.first, p.last, p.age)
    p.last = "nair"
    fmt.Println(p.first, p.last, p.age)
}

func (p *person) printNameB() {
    fmt.Println(p.first, p.last, p.age)
    fmt.Println((*p).first, (*p).last, (*p).age)
    
    p.last = "ramachandran nair"
    fmt.Println(p.first, p.last, p.age)
}
