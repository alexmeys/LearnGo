package main

import (
	"fmt"
)

func main() {
	var y int = 112

	// Y printed
	fmt.Println(y)

	// Type Y printed
	fmt.Printf("%T\n\n", y)

	// binary, hex, hex with 0x and default value
	// godoc.org/fmt (# = other flags)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	fmt.Printf("%v\n", y)

	// one liner
	fmt.Printf("\n%b\t%x\t%#x\t%v\n", y, y, y, y)

	// print one liner with string
	s := fmt.Sprintf("\n%b\t%x\t%#x\t%v\n", y, y, y, y)

	fmt.Println(s)
	fmt.Println("\nThat's all folks")
}
