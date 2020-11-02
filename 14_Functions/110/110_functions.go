package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(xi...)
	fmt.Println("All:", s)

	s2 := even(sum, xi...)
	fmt.Println("Even:", s2)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

//callback
func even(f func(xi ...int) int, ii ...int) int {
	var yi []int
	for _, v := range ii {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
