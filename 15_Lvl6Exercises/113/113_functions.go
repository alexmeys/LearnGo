package main

import "fmt"

func main() {
	x := foo()
	y, v := bar()
	fmt.Println(x)
	fmt.Println(y, v)
}
func foo() int {
	var x int = 20
	return x
}
func bar() (int, string) {
	x := 10
	s := "Hello there"
	return x, s
}
