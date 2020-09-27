package main

import "fmt"

func main() {
    p1 := person{age: 35, lname: "rp", fname: "adersh"}
    p1.printNameA()
    p1.printNameB()
}

type person struct {
    fname string
    lname string
    age   int
}

func (p person) printNameA() {
    fmt.Println(p.fname, p.lname, p.age)
    p.lname = "nair"
    fmt.Println(p.fname, p.lname, p.age)
}

func (p *person) printNameB() {
    fmt.Println(p.fname, p.lname, p.age)
    fmt.Println((*p).fname, (*p).lname, (*p).age)
    
    p.lname = "ramachandran nair"
    fmt.Println(p.fname, p.lname, p.age)
}
