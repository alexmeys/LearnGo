package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	fmt.Printf("Val x:%T\n", x)
	fmt.Printf("Val x:%v\n", x)
	fmt.Printf("Val y:%T\n", y)
	fmt.Printf("Val y:%v\n", y)
	fmt.Printf("Val z:%T\n", z)
	fmt.Printf("Val z:%v\n", z)
}
