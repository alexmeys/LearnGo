package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Recap Channels - basics")
	// first part, basic channel put on and get off channel
	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	fmt.Println(<-ch)
	fmt.Println("-------")
	// buffered channel
	ch2 := make(chan int, 2)
	go alef(ch2)
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
	fmt.Println("-------")
	// range over loop
	ch3 := make(chan int)
	go be(ch3)
	for v := range ch3 {
		fmt.Println(v)
	}
	fmt.Println("-------")
	// receiver / sender separate
	ch4 := make(chan int)
	ch5 := make(chan int)
	ch6 := make(chan int)

	go pe(ch4, ch5, ch6)
	te(ch4, ch5, ch6)
	fmt.Println("-------")
	// fan in
	ch7 := make(chan int)
	ch8 := make(chan int)
	ch9 := make(chan int)

	go se(ch7, ch8)
	go je(ch7, ch8, ch9)
	for v := range ch9 {
		fmt.Println(v)
	}
	fmt.Println("-------")
	// fan out
	ch10 := make(chan int)
	ch11 := make(chan int)

	go che(ch10)
	go he(ch10, ch11)
	for v := range ch11 {
		fmt.Println(v)
	}

	fmt.Println("-------")
	// Done
	fmt.Println("That's all folks!")
}

func alef(ch chan int) {
	ch <- 43
	ch <- 44
}

func be(ch chan int) {
	for i := 45; i < 51; i++ {
		ch <- i
	}
	close(ch)
}

func pe(c1, c2, c3 chan<- int) {
	for i := 50; i < 56; i++ {
		if i%2 == 0 {
			c1 <- i
		} else {
			c2 <- i
		}
	}
	//close(c1)
	//close(c2)
	close(c3)
}

func te(c1, c2, c3 <-chan int) {
	for {
		select {
		case v := <-c1:
			fmt.Println(v)
		case v := <-c2:
			fmt.Println(v)
		case <-c3:
			return
		}

	}

}

func se(c1, c2 chan<- int) {
	for i := 55; i < 61; i++ {
		if i%2 == 0 {
			c1 <- i
		} else {
			c2 <- i
		}
	}
	close(c1)
	close(c2)
}

func je(c1, c2 <-chan int, c3 chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range c1 {
			c3 <- v
		}
		wg.Done()
	}()
	go func() {
		for v := range c2 {
			c3 <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(c3)
}

func che(c1 chan int) {
	for i := 61; i < 66; i++ {
		c1 <- i
	}
	close(c1)
}

func he(c1, c2 chan int) {
	var mu sync.Mutex
	var wg sync.WaitGroup
	total := 61
	for v := range c1 {
		wg.Add(1)
		go func(v2 int) {
			mu.Lock()
			c2 <- ghe(total)
			total++
			mu.Unlock()
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func ghe(c int) int {
	return c
}
