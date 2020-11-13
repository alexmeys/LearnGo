package main

import "fmt"

func main() {
	x := 0
	fmt.Println("x start", &x)
	fmt.Println("x start", x)
	foo(&x)
	fmt.Println("x after", &x)
	fmt.Println("x after", x)
}

func foo(y *int) {
	fmt.Println("y start", y)
	fmt.Println("y start", *y)
	*y = 43
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}
