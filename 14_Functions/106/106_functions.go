package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("xxxxxxxx", h.(person).first)
	case secretAgent:
		fmt.Println("yyyyyyyy", h.(secretAgent).first)
	}
	fmt.Println("I'm human", h)
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "Jeff",
			last:  "Jefferson",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			first: "John",
			last:  "Johnson",
		},
		ltk: true,
	}
	p1 := person{
		first: "Mr",
		last:  "Nobody",
	}

	bar(sa1)
	bar(sa2)
	bar(p1)
}
func (s secretAgent) speak() {
	fmt.Println(s.first, s.last, "-The SecretAgent speak")
}
func (p person) speak() {
	fmt.Println(p.first, p.last, "-The person speak")
}
