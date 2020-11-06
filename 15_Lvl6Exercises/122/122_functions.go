package main

import "fmt"

func main() {
	x := 10

	{
		a := 20
		z := foo()
		fmt.Println(x + a + z)
	}

}
func foo() int {
	c := 100
	return c
}
