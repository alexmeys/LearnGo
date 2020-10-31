package main

import "fmt"

func main() {
	ms := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(ms[0:5])
	fmt.Println(ms[5:])
	fmt.Println(ms[2:7])
	fmt.Println(ms[1:6])

}
