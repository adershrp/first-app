package main

import (
	"fmt"
)

func main() {
	a := "Start"
	defer fmt.Println(a)
	a = "End"
	defer fmt.Println(a) //defer will store the value of a at the time of calling this defer.
}
