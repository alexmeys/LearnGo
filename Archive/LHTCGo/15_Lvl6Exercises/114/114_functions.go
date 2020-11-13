package main

import "fmt"

func main() {
	y := []int{2, 4, 6, 8}

	z := foo(y...)
	fmt.Println(z)

	w := bar(y)
	fmt.Println(w)
}

func foo(y ...int) int {
	var total int = 0
	for _, v := range y {
		total += v
	}
	return total
}

func bar(w []int) int {
	var total int = 0
	for _, v := range w {
		total += v
	}
	return total
}
