package main

import "fmt"

func main() {
	type person struct {
		fname string
		lname string
		age   uint8
	}

	type col struct {
		person
		years uint8
	}

	p1 := person{
		fname: "Peter",
		lname: "Peterson",
		age:   40,
	}

	p2 := person{
		fname: "Jef",
		lname: "Jefferson",
		age:   40,
	}

	p3 := col{
		years: 10,
		person: person{
			fname: "John",
			lname: "Johnson",
			age:   40,
		},
	}

	fmt.Println(p1.fname, p1.lname, p1.age)
	fmt.Println(p2.fname, p2.lname, p2.age)
	// Prompted field p3.years
	fmt.Println(p3.fname, p3.lname, p3.age, p3.years)
}
