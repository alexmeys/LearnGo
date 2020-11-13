package main

import (
	"fmt"
)

func main() {

	// 2 in binary
	x := 2
	fmt.Printf("%d\t\t%b\n", x, x)

	// bit shifting
	// shifts the 1 to left and added a 0
	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

}
