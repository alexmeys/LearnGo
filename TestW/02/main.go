package main

import (
	"fmt"
)

func main() {

	ch := make(chan string)

	xs := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
	}

	transfer(xs, ch)

	for i := 0; i < len(xs); i++ {
		fmt.Printf("%v: %v\n", xs[i], <-ch)
	}

}

func transfer(xs []string, ch chan string) {
	fmt.Println("Test channels")
	// random order for int's.
	for _, v := range xs {
		go func(v string) {
			ch <- v
		}(v)
	}
}
