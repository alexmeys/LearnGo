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
		favIce: []string{"Strawberry", "hazelnut"},
	}

	m := map[string]person{
		p1.lname: p1,
		p2.lname: p2,
	}

	for k, v := range m {
		fmt.Println(k)
		for _, v2 := range v.favIce {
			fmt.Printf("\t%v\n", v2)
		}
	}

}
