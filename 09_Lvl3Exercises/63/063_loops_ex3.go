package main

import "fmt"

func main() {

	var total uint8 = 0
	var i uint32 = 2020
	for i > 1920 {
		fmt.Printf("%d\n", i)
		i--
		total++
	}
	fmt.Println("You have been alive for ", total, "years...")
}
