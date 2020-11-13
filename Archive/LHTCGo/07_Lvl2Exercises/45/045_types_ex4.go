package main

import "fmt"

func main() {
	var a rune = 120
	fmt.Printf("%v\t%b\t\t%#x\n", a, a, a)

	c := a << 1
	fmt.Printf("%v\t%b\t%#x\n", c, c, c)

}
