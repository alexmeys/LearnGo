package main

import "fmt"

func main() {
	var x bool
	fmt.Println(x)
	x = true
	fmt.Println(x)

	// bool usage
	a := 7
	b := 42

	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a <= b)
	fmt.Println(a >= b)

}
