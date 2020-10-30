package main

import (
	"fmt"
)

func main() {
	switch {
	default:
		fmt.Println("This is default")
	case false:
		fmt.Println("False")
	case (2 == 4):
		fmt.Println("noprint")
	case (3 == 2):
		fmt.Println("Printed")
		// continue the loop normally, however it prints everything below
		// dont use that
		fallthrough
	case (5 == 4):
		fmt.Println("NotPrinted?")
	}
}
