package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go foo(ch)

	for v := range ch {
		fmt.Println(v)
	}
}

func foo(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
