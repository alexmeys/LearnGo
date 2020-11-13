package main

import "fmt"

func main() {
	switch {
	default:
		fmt.Println("In default")
	case true:
		fmt.Println("In true")
	case false:
		fmt.Println("In false")
	}
}
