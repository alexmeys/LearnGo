package main

import "fmt"

func main() {

	// Anonymous struct
	// has no name, and immediate declaration in p1

	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   24,
	}

	fmt.Println(p1)

}
