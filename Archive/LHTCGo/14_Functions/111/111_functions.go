package main

import "fmt"

func main() {
	a := incremntor()
	b := incremntor()
	// different identifier a - b scope of x example
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func incremntor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
