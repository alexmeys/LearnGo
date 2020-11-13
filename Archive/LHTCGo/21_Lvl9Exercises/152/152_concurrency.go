package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Data race fix with atomic
func main() {
	var wg sync.WaitGroup
	var count int64
	const gs = 100

	wg.Add(gs)
	for i := 1; i <= gs; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			fmt.Println(atomic.LoadInt64(&count))
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Total:", count)
}
