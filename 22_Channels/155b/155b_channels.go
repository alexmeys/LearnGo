package main

import "fmt"

func main() {
	// buffer channel (try to stay away from it but know it)
	// unbuffered is better
	// allows certani values to sit in that channel
	// here 1 values stored on the channel
	c := make(chan int, 2)

	// put on the channel with 42
	c <- 42
	c <- 43

	// print it from the channel.
	fmt.Println(<-c)
	fmt.Println(<-c)
}
