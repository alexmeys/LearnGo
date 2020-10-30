package main

import "fmt"

func main() {
	var tekst string = "dit is een lange tekst"
	var part string = "dit"
	var full string = "dit is een lange tekst"

	switch tekst {
	default:
		fmt.Println("Default triggered")
	case part:
		fmt.Println("Found it in partly text")
	case full:
		fmt.Println("Found it in full string")

	}
	// need to check in future lessons if pattern matching
	//can be applied in switch to get info out of text
	fmt.Println("")
}
