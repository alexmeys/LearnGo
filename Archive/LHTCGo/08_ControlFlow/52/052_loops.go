package main

import "fmt"

func main() {
	// This can be used for while (absent in Go)
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("Done")

	y := 1
	for {
		if y > 9 {
			break
		}
		fmt.Println(y)
		y++
	}
	fmt.Println("Done 2")
}
