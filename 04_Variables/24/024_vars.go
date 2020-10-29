package main

import (
	"fmt"
)

func main() {
	//Declare a variable to be of certain TYPE
	// Assign a value of that type to the variable
	// you cannot assign another type to it

	// Zero values
	var y string
	var z int

	fmt.Println("Y:", y)
	fmt.Printf("Type: %T\n", y)

	fmt.Println("Z:", z)
	fmt.Printf("Type: %T\n", z)

	fmt.Println("That's all folks")

}
