package main

import "fmt"

func main() {
	const (
		// a = typed, smaller faster but fixed type
		// b = untyped gives flexibility vs size
		a uint8 = 10
		b       = 254
	)

	fmt.Println(a)
	fmt.Println(b)

}
