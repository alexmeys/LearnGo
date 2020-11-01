package main

import "fmt"

func main() {
	type person struct {
		first string
		last  string
		age   uint8
	}

	// embedding types
	type secretAgent struct {
		person
		ltk bool
		// name collision, not best practice
		first string
	}
	p1 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Println(p1.last)

	p2 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		first: "name collision",
		ltk:   true,
	}

	fmt.Println(p1.last, p1.first, p1.age)
	fmt.Println(p2.last, p2.person.first, p2.age, p2.ltk, "Collsion:", p2.first)
}
