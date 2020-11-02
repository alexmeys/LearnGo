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

func main() {

	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			"Jef",
			"Jefferson",
		},
		ltk: true,
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
}

func (s secretAgent) speak() {
	fmt.Println("I'm", s.first, s.last)
}
