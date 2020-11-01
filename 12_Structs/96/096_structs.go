package main

import "fmt"

func main() {
	// var - identifier - type
	var x int
	// var - identifier - type
	type person struct {
		first string
		last  string
	}
	// var - identifier - type
	const bar int = 42

	p1 := person{
		first: "ABC",
		last:  "DEF",
	}

	// nonono
	type foo int
	var y foo = 444

	var z = y

	// does not work, as bar is const type
	// w is declared as foo type, so does not work
	// w without foo, var w = bar would work
	// var w foo = bar

	fmt.Println(x)
	fmt.Println(bar)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(p1.first)

}
