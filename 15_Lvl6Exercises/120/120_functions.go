package main

import "fmt"

func main() {
	x := fret()
	i := x()
	fmt.Println(i)
}

func fret() func() int {

	return func() int { return 20 }
}
