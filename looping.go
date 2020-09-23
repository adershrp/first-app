package main

import (
	"fmt"
)

func main() {

	//for init;condition;post
	//normal for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	//two variables. - cannot use expression for incrementing
	for i, j := 1, 1; i <= 12 && j < 10; i, j = i+1, i*2 {
		fmt.Println(i, j)
	}

	//i initialized out side for-loop helps in getting the value of i outside forloop context
	i := 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("value of i outside for loop", i)
	// dowhile and while loop have to be done this way
	j := 0
	for j < 5 {
		j++
		fmt.Println("do-while || while", j)
	}

	k := 0
	//infinite loop
	for {
		k++
		fmt.Println("Infinite loop", k)
		if k == 5 { // loop breaking condition
			break
		}
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // stops the flow and return control back to for-loop
		}
		fmt.Println("Print where above condition fails", i)
	}
Loop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			k := i * j
			if k > 50 {
				break Loop // break the loop using labels
			}
			fmt.Println(i, " * ", j, " = ", k)
		}
	}
	// iterating array / slice
	ints := []int{1, 2, 3, 4, 5, 6}
	for key, val := range ints {
		fmt.Println(key, val)
	}
	// map iteration
	maps := map[string]int{
		"Sun": 1,
		"Mon": 2,
		"Tue": 3,
		"Wed": 4,
		"Thu": 5,
		"Fri": 6,
		"Sat": 7,
	}
	for key, val := range maps {
		fmt.Println(key, val)
	}
}
