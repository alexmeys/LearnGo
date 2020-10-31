package main

import "fmt"

func main() {
	x := []int{0, 11, 22, 33, 44, 55, 66, 99}
	fmt.Println(x)

	// Delete from slice the first and last element
	// basically make a new slice x from x1:7 (in 0..7 items),
	// the last digit is uppper limit and will be discarded, want more? add 1
	x = append(x[1:7])
	fmt.Println(x)

	y := []int{4, 5, 7, 8, 42, 77, 88, 99, 1014, 234, 456, 678, 987}
	fmt.Println(y)

	// This will remove digits 2-3
	// first argument is the []int slice, then we extract the second slice with ....
	// and this will convert it to int and append it to the first slice again
	y = append(y[:2], y[4:]...)

	fmt.Println(y)

}
