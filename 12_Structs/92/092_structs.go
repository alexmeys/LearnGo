package main

import "fmt"

func main() {

	// composite data structure /
	//complex datastructre  /
	// aggregate data type = struct

	type person struct {
		first string
		last  string
		age   uint8
	}

	p1 := person{
		first: "James",
		last:  "Bond",
		age:   37,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   32,
	}

	fmt.Println(p1, p2)
	fmt.Println(p1.first)
	fmt.Println(p2.last)
	fmt.Println(p1.age)

}
