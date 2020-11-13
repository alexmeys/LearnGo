package main

import "fmt"

func main() {
	defer foo()
	bar()
	fmt.Println("...more code...")

	//here runs defer (just before main exit)
}

func foo() {
	fmt.Println("Foo")
}

func bar() {
	fmt.Println("Bar")
}
