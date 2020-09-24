package main

import (
	"fmt"
)

//struct -- allow to create datastructure of different data type.
// help to aggregate different datattype to create a custom data type
type person struct {
	fname string
	lname string
	age   int
}

//embeded struct - not inheritence ... composite relation
type parent struct {
	person
	isParent bool
}

func main() {
	p1 := person{
		fname: "Adersh", lname: "R P", age: 35,
	}
	p2 := person{
		fname: "Divya", lname: "Challiyil", age: 28,
	}
	p3 := person{
		fname: "Prayag", lname: "Nair", age: 3,
	}

	father := parent{
		person: p1, isParent: true,
	}

	mother := parent{
		person: p2, isParent: true,
	}

	child := parent{
		person: p3, isParent: false,
	}

	//Anonymous Struct
	employe := struct {
		fname string
		lname string
		age   int
	}{
		fname: "Adersh", lname: "Nair", age: 35,
	}

	fmt.Println("Anonymous Struct", employe.fname, employe.lname, employe.age)

	fmt.Println(father.fname, father.lname, father.age)
	fmt.Println(mother.fname, mother.lname, mother.age)
	fmt.Println(child.fname, child.lname, child.age)

}
