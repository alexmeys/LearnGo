package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()
	fmt.Println(<-c)

	ch := make(chan int, 1)
	ch <- 43
	fmt.Println(<-ch)

}
