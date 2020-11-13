package main

import (
	"fmt"
	"runtime"
	"sync"
)

// func init is to initialise code
// func init() {}

// add a wait group to global scope
var wg sync.WaitGroup

func main() {

	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	// tell it to wait for 1 thing
	wg.Add(1)
	// spin off it's own thing
	go foo()

	bar()
	// Concurrent 2 goroutines will be printed at this time (main - go foo())
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	// Wait for it
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	// Done signal send back to main and will decrement the wg.Add and wg.Wait() will finish
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
