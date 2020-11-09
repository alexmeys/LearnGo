package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// send
	go foo(c)

	// receive
	// remove go in front, so it's blocking until foo is done
	bar(c)

	fmt.Println("...Exiting...")
}

// send
func foo(c chan<- int) {
	c <- 42
}

// receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}
