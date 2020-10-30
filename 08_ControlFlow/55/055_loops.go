package main

import "fmt"

func main() {
	if false {
		fmt.Println("All Done with 001 false")
	}
	if true {
		fmt.Println("All Done with 002 true")
	}
	if !false {
		fmt.Println("All Done with 003 !false")
	}
	if !true {
		fmt.Println("All Done with 004 !true")
	}
	fmt.Printf("All Done - Part 1\n\n")

	// limit scope of x to this if loop
	if x := 42; x == 42 {
		fmt.Println("It ran 2 expressions in 1 line")
	}
	// fmt.Println(x)
}
