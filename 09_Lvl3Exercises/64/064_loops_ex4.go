package main

import "fmt"

func main() {
	i := 2020
	for {
		i--
		if i == 1920 {
			fmt.Println("Finally arrived at your birth year, the great year", i)
			break
		} else {
			fmt.Println(i)
		}

	}
}
