package main

import "fmt"

func main() {
	n := factorial(4)
	fmt.Println(n)

	p := loopFact(4)
	fmt.Println(p)

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loopFact(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}

// 4 * 3 * 2 * 2 * 1 * (0)->1
