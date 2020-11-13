package main

import "fmt"

func main() {
	x := []int{9, 8, 22, 33, 44}
	fmt.Println(x)
	fmt.Println(len(x))

	for i, v := range x {
		fmt.Println(i, v)
	}

}
