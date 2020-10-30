package main

import "fmt"

func main() {
	z := 10
	y := 20

	a := (z == y)
	b := (z <= y)
	c := (z >= y)
	d := (z != y)
	e := (z < y)
	f := (z > y)

	fmt.Println(a, b, c, d, e, f)

}
