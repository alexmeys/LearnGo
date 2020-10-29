package main

import "fmt"

func main() {
	s := "Alex"

	// slices and converts string to decimal values
	var sConv = []byte(s)

	fmt.Printf("The sConv is of type: %T\n\n", sConv)

	fmt.Println([]byte(s))

	// manuel hexing without 0x
	fmt.Printf("%x\t", sConv[0])
	fmt.Printf("%x\t", sConv[1])
	fmt.Printf("%x\t", sConv[2])
	fmt.Printf("%x\n", sConv[3])

	// manuel utf8
	fmt.Printf("%U\t", sConv[0])
	fmt.Printf("%U\t", sConv[1])
	fmt.Printf("%U\t", sConv[2])
	fmt.Printf("%U\n", sConv[3])

	// binary
	fmt.Printf("%b\t", sConv[0])
	fmt.Printf("%b\t", sConv[1])
	fmt.Printf("%b\t", sConv[2])
	fmt.Printf("%b\n", sConv[3])
}
