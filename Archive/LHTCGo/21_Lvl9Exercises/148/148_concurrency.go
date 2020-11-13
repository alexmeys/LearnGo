package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	fmt.Println("Go's:", runtime.NumGoroutine())
	go foo()
	go func() {
		for i := 9; i <= 20; i++ {
			fmt.Println("Main:", i)
		}
		wg.Done()
	}()
	fmt.Println("Go's:", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("All Done")
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo:", i)
	}
	wg.Done()
}
