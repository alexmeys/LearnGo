package main

import "fmt"

func main() {
	func() { fmt.Println("Anon") }()

	func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()

	z := func() int {
		x := 20
		return x
	}()

	y := func(z int) int {
		a := z * 2
		return a
	}(z)

	fmt.Println("Z:", z)
	fmt.Println("Y:", y)
}
