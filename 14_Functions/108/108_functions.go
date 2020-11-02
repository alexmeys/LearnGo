package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("My first func expression")
	}
	f()

	f2 := func(x int) {
		fmt.Println("This year is:", x)
	}
	f2(2020)
}
