package main

import "fmt"

func main() {
	// true
	//fmt.Println(true && true)
	//fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!false)
	fmt.Println("")
	// false
	fmt.Println(true && false)
	fmt.Println(!true)
	fmt.Println("")
	// deeper dive
	fmt.Println(!true && false)
	fmt.Println(!false && true)
	//fmt.Println(!false || !false)
	//fmt.Println(!true || !true)
	fmt.Println(!true || false)
	fmt.Println(!false || true)
}
