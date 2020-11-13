package main

import (
	"fmt"
)

func main() {
	var a int = 24

	type hotdog int
	var b hotdog = 32

	fmt.Println(a, b)
	fmt.Printf("%T", a)
	fmt.Printf("%T\n\n", b)

	var c int = int(b)

	fmt.Println(a, c)
	fmt.Printf("%T", a)
	fmt.Printf("%T\n\n", c)

	fmt.Println("That's all folks")
}
