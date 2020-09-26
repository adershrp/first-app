package main

import (
    "fmt"
)

// struct - aggregate of different data type
type person struct {
    fname, lname string
    age          int
}

// embeded struct -> composite relation
type parent struct {
    person
    isParent bool
    company  string
}

// embeded struct
type children struct {
    person
    isParent bool
    school   string
}

// method -- adding context to the struct
func (p person) speak() string {
    return fmt.Sprint(p.fname, " ", p.lname)
}

// method -- adding context to the struct
func (p parent) speak() string {
    return fmt.Sprint("Parent Name is ", p.person.speak(), " Working at ", p.company)
}

// method -- adding context to the struct
func (p children) speak() string {
    return fmt.Sprint("Kid Name is ", p.person.speak(), " Studying at ", p.school)
}

// main method
func main() {
    // father struct
    father := parent{
        person:   person{fname: "Adersh", lname: "RP", age: 35},
        isParent: true,
        company:  "ibm",
    }
    mother := parent{
        person:   person{fname: "Divya", lname: "Challiyil", age: 28},
        isParent: true,
        company:  "siemens",
    }
    kid := children{
        person:   person{fname: "Prayag", lname: "Nair", age: 3},
        isParent: false,
        school:   "Kumarans",
    }
    defer fmt.Println(father.speak())
    defer fmt.Println(mother.speak())
    fmt.Println(kid.speak())
    
    // fmt.Println(father, mother, kid)
    
}
