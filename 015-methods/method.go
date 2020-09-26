package main

import "fmt"

func main() {
    p1 := person{
        fname: "adersh", lname: "rp", age: 35,
    }
    p2 := newPerson("divya", "challiyil", 28)
    
    fmt.Println(p1.speak())
    fmt.Println(p2.speak())
}

type person struct {
    fname, lname string
    age          int
}

// constructor
func newPerson(fname string, lname string, age int) *person {
    return &person{fname: fname, lname: lname, age: age}
}

func (p person) speak() string {
    return fmt.Sprint("My Name is ", p.fname, " ", p.lname, " and aged ", p.age)
}
