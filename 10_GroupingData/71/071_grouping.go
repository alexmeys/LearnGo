package main

import (
	"fmt"
)

func main() {

	// Array of int's 1
	var x [5]int

	fmt.Println(x)

	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))

	var y [2][8]int

	//fmt.Println(y)
	y[0][0] = 2
	y[0][1] = 3
	y[0][2] = 4
	y[0][3] = 5
	y[0][4] = 6
	y[0][5] = 4
	y[0][6] = 3
	y[0][7] = 2

	for i := 0; i < 8; i++ {
		y[1][i] = 1
	}

	fmt.Println(y)

}
