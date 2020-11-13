package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)

	x := bar()

	i := x()
	fmt.Println(i)

}

// return string
func foo() string {
	s := "Hello"
	return s
}

// return a function of type int
func bar() func() int {
	return func() int {
		return 2020
	}
}
