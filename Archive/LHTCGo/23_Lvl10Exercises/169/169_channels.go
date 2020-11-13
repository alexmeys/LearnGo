package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func(b chan<- int) {
		for i := 0; i < 100; i++ {
			b <- i
		}
		close(b)
	}(c)

	func(a <-chan int) {
		for v := range a {
			fmt.Println(v)
		}

	}(c)
}
