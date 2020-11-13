package main

import "fmt"

func main() {
	A := 20
	B := 10
	if A == B {
		fmt.Println("A == B")
	} else if A < B {
		fmt.Println("A Is smaller than B")
	} else if A > B {
		fmt.Println("A is bigger than B")
	} else {
		fmt.Println("A != B, no idea how u got here")
	}
}
