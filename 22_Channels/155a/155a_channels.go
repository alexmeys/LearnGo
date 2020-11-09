package main

import (
	"fmt"
)

func main() {
	// create a channel
	c := make(chan int)

	// start a goroutine
	go func() {
		// put 42 on the channel
		c <- 42
	}()

	// wait until block is lifted, print c off that channel
	fmt.Println(<-c)

}
