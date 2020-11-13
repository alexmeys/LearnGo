package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("Welcome to concurrency")
	wg.Add(1)
	go foo()
	bar()
	wg.Wait()
}

func foo() {
	for i := 10; i < 20; i++ {
		fmt.Println(i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
