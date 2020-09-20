package main

import (
	"fmt"
)

func main() {
	number := 8
	switch number {
	case 1, 3, 5, 7, 9:
		fmt.Println("Odd Number")
		fmt.Printf("%T\t%v\n", number, number)
	case 2, 4, 6, 8, 10:
		fmt.Println("Even Number")
		fmt.Printf("%T\t%v\n", number, number)
	default:
		fmt.Println("Another Number")
	}

	textSample := "Monday"
	switch textSample {
	case "Sunday", "Saturday":
		fmt.Println("Weekend")
	case "Monday", "Tuesday":
		fmt.Println("Weekday")

	}
	switch i := 1 + 4; i {
	case 1, 3, 5, 7, 9:
		fmt.Println("Odd Number", i)
	case 2, 4, 6, 8, 10:
		fmt.Println("Even Number", i)
	}
}
