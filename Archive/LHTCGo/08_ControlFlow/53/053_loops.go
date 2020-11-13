package main

import "fmt"

func main() {
	// Get it started from 0
	x := -1
	for {
		x++
		if x >= 100 {
			// breaks the (for) loop
			break
		}
		if x%2 != 0 {
			// continue goes back to top execution in for
			continue
		}
		// print the number if it's %2
		fmt.Println(x)

	}
	// finish
	fmt.Println("The end")
}
