package main

import (
	"fmt"
)

func main() {

	// creates a channel "c"
	c := make(chan int)
	// puts 42 on the channel - and locks
	// channels block
	c <- 42
	// try and print the value - does not work.
	fmt.Println(<-c)

}
