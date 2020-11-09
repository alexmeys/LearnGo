package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // run at the end of code

	for v := range gen(ctx) {
		fmt.Println(v)
		if v == 5 {
			break
		}
	}
}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // returns , not to leak the goroutine
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}
