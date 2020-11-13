package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("Outer: %d\n", i)
		for j := 0; j < 2; j++ {
			fmt.Printf("\t\tInner: %d\n", j)
		}
	}
}
