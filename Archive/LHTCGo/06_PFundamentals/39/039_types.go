package main

import "fmt"

func main() {

	// multipe constants without needing to type const on every line
	// typed and untyped to come...
	const (
		d int     = 20
		e float64 = 21.22
		f uint8   = 23
	)

	const a = 42
	const b = 42.789
	const c = "Alex"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
