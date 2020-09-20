package main

import (
	"fmt"
)

type Employee struct {
	id   int32
	name string
}

type SoftwareEngineer struct {
	Employee
	designation string
	language    []string
}

func main() {

	emp1 := SoftwareEngineer{}
	emp1.id = 1
	emp1.designation = "Architect"
	emp1.language = []string{"Java", "JavaScript"}
	emp1.name = "adersh"

	emp2 := SoftwareEngineer{}
	emp2.id = 2
	emp2.name = "divya"
	emp2.designation = "Senior Engineer"
	emp2.language = []string{"ruby"}

	fmt.Println(emp1)
	fmt.Println(emp2)
	fmt.Println(emp2.name)

}
