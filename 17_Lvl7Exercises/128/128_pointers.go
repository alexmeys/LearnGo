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
	changeme(&p1)
	fmt.Println(p1.name)
	changeme2(&p1)
	fmt.Println(p1.name)
}

func changeme(p *person) {
	p.name = "Johnson"
}
func changeme2(p *person) {
	// use (*p) for dereference the pointer here!
	(*p).name = "kevinson"
}
