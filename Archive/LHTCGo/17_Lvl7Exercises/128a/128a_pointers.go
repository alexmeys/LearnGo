package main

import "fmt"

type person struct {
	name string
}

func main() {

	p1 := person{
		"Jefferson",
	}
	fmt.Println(p1.name)
	p1 = changeme(p1)
	fmt.Println(p1.name)
	p1 = changeme2(p1)
	fmt.Println(p1.name)
}

func changeme(p person) person {
	p.name = "Johnson"
	return p
}
func changeme2(p person) person {
	p.name = "kevinson"
	return p
}
