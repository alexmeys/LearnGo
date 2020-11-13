package main

import (
	"fmt"
)

func main() {

	// basic type int, printed
	var a int = 42
	fmt.Println("\n", a)
	fmt.Printf("%T\n\n", a)

	// created type 'hotdog' of int
	// you cannot cast custom type to int, even they are both int's
	type hotdog int

	var b hotdog = 43
	fmt.Println(b)
	fmt.Printf("%T\n\n", b)

	fmt.Print("That's all folks\n")
}
