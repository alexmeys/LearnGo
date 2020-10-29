package main

import (
	"fmt"
)

func main() {

	var y = 42
	fmt.Println(y)

	// Printf %T, prints the type of a variable.
	fmt.Printf("%T\n", y)

	// This will not work:
	// y = "Some string content"
	// Y is declared as an INT type, you cannot cross types

	var z string = "this is a string"
	fmt.Println(z)

}
