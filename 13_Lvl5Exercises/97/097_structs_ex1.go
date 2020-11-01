package main

import "fmt"

func main() {
	type person struct {
		fname  string
		lname  string
		favIce []string
	}

	p1 := person{
		fname:  "Jef",
		lname:  "Jefferson",
		favIce: []string{"Vanilla", "Chocolate"},
	}

	p2 := person{
		fname:  "John",
		lname:  "Johnson",
		favIce: []string{"Strawberry", "Hazelnut"},
	}

	fmt.Println(p1.fname, "favorite ice cream:")
	for _, v := range p1.favIce {
		fmt.Printf("\t%v\n", v)
	}

	fmt.Println(p2.fname, "favorite ice cream:")
	for _, v := range p2.favIce {
		fmt.Printf("\t%v\n", v)
	}
}
