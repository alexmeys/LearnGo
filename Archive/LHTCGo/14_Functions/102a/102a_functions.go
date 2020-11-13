package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	x := foo(xi...)
	fmt.Println("If we count your numbers, the total is:", x)

	// calling defer on 2nd print statement will print it very last before exiting main.
	y := even(xi...)
	defer fmt.Println("If we only count up the even numbers we get:", y)

	z := uneven(xi...)
	fmt.Println("If we only count up the uneven number s we get:", z)
}

func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	sum := 0
	for _, v := range x {
		sum = (sum + v)
	}

	return sum

}

// variadic parameter (meaning 0 or more elements)
// Here 'x' takes the int's and put it in an array / slice of int's.
// below we see for range to reach those elements.

func even(x ...int) int {
	var sum int
	for _, v := range x {
		if v%2 == 0 {
			sum = (sum + v)
		}
	}
	return sum

}

func uneven(x ...int) int {
	var sum int
	for _, v := range x {
		if v%2 != 0 {
			sum = (sum + v)
		}
	}
	return sum
}
