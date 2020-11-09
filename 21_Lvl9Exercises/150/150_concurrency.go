package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Data race
func main() {
	var wg sync.WaitGroup
	count := 0
	const gs = 100

	wg.Add(gs)
	for i := 1; i <= gs; i++ {
		go func() {
			v := count
			runtime.Gosched()
			v++
			count = v
			fmt.Println(count)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Total:", count)
}
