package main

import "fmt"

func main() {
	fmt.Println("2+3=", mySum(2, 3))
	fmt.Println("2+8=", mySum(2, 8))
	fmt.Println("2+12+8=", mySum(2, 12, 7))
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
