package main

import "fmt"

func main() {
	nos := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sum(nos, odd))
	fmt.Println(sum(nos, even))

}

func sum(xi []int, f func(nos []int) []int) int {
	var total int = 0
	for _, v := range f(xi) {
		total += v
	}
	return total
}

//return odd number
func odd(xi []int) []int {
	var odd []int
	for _, v := range xi {
		if v%2 != 0 {
			odd = append(odd, v)
		}
	}
	return odd
}

//return even numbers
func even(xi []int) []int {
	var even []int
	for _, v := range xi {
		if v%2 == 0 {
			even = append(even, v)
		}
	}
	return even
}
