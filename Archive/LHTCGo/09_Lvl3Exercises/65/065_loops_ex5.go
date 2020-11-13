package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Println("The remainer for", i, ":", i%4)
	}
}
