package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Data race fix with mutexes
func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	count := 0
	const gs = 100

	wg.Add(gs)
	for i := 1; i <= gs; i++ {
		go func() {
			mu.Lock()
			v := count
			v++
			count = v
			fmt.Println("NUM:", count)
			fmt.Println("GO:", runtime.NumGoroutine())
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Total:", count)
}
