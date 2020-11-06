package main

import "fmt"

func main() {
	var x int = 20
	var y int = 10
	func() { fmt.Println("Print the int", x) }()

	func(x int) { fmt.Println("20 added by 20 =", x+20) }(x)

	fmt.Println("Result of y:", y)

	z := func(y int) int { return y }
	y = z(2020)

	fmt.Println("New Result of y:", y)

}
