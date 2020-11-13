package main

import "fmt"

func main() {
	ms := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

	for k, v := range ms {
		fmt.Println("Key:", k, "Value:", v)
	}
	fmt.Printf("Type slice: %T\n", ms)
}
