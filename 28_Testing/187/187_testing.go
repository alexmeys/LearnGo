package main

import "fmt"

func main() {
	fmt.Println("2+3=", mySum(2, 3))
	fmt.Println("4+3=", mySum(4, 3))
	fmt.Println("6+3=", mySum(6, 3))
}

func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum + 1
}
