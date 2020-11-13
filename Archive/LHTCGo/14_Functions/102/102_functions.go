package main

import "fmt"

func main() {
	x := foo(6, 4, 1, 5, 8, 9, 7)

	fmt.Println("From main:", x)
}

// received individual ints -> slice of int inside function
func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("For item in indexp pos:", i, "we are now adding:", v, "to the total, this makes:", sum)
	}
	fmt.Println("Grand total:", sum)
	return sum
}
