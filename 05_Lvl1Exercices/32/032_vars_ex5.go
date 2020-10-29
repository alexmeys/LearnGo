package main

import (
	"fmt"
)

type alex int

var x alex
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 112
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
