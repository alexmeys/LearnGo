package main

import (
	"fmt"
)

// unsuccesful buffer
func main() {
	c := make(chan int, 1)
	c <- 42
	// blocked by trying to put the value on the channel.
	c <- 43

	fmt.Println(<-c)
}
