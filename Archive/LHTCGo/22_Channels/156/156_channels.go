package main

import (
	"fmt"
)

func main() {

	c := make(chan int, 1) // bi-directional channel
	cr := make(<-chan int) // receive channel only
	cs := make(chan<- int) // send channel only
	fmt.Println("-----")
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)
}
