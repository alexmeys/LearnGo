package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 43, 17, 987, 14, 123, 21, 1, 4, 2}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry"}

	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println(xi)
	fmt.Println(xs)
}
