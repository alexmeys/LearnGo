package main

import "fmt"

func main() {
	const (
		a = iota
		b = iota
		c = iota
	)
	// This would reset the count to 0 of the previous iota, otherwise it would increment
	// const word reset the count
	const (
		d = iota
		e
		f
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

}
