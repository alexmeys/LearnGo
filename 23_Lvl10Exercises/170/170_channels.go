package main

import (
	"fmt"
)

func main() {
	const goroutines = 10
	ch1 := make(chan int)
	q := make(chan int)

	func() {
		go func() {
			for i := 0; i < goroutines; i++ {
				for j := 0; j < goroutines; j++ {
					ch1 <- (j + (i * 10))
				}
			}
			q <- 1
		}()
	}()

	func() {
		for {
			select {
			case v := <-ch1:
				fmt.Println(v)
			case <-q:
				return
			}
		}
	}()

}
