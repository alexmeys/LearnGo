package main

import (
	"fmt"
)

// outside scope
// x := 42 here would not work
// var y = 43, works here and can be access in every function below.

func main() {
	// Declares a variable and assign a value (of a certain type)
	// Short declaration operator
	x := 42
	fmt.Println("Hello", x)

	var y = 43
	fmt.Println(y)

	// declares a variable z of type int
	// var z int without assignment can work, but = false, 0, nil,...
	var z int = 44
	fmt.Println(z)

	var a = 20
	fmt.Println(a)

}
