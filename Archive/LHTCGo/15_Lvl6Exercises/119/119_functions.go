package main

import "fmt"

func main() {
	s := foo()

	fmt.Println(s)
}

func foo() string {
	return "The string"
}
