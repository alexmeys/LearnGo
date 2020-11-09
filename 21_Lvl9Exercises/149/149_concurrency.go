package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
	age   uint8
}

type human interface {
	speak()
}

func main() {
	fmt.Println("Start our program")
	fmt.Println("")
	h1 := person{
		fname: "Jef",
		lname: "Jefferson",
		age:   20,
	}
	saySomething(&h1)
	fmt.Println("")
	fmt.Println("End our program")
}

func (p *person) speak() {
	fmt.Println("lets try to speak")
	fmt.Println("")
	fmt.Println("My name is:", p.fname, p.lname)
	fmt.Println("I'm", p.age, "years old.")
}

func saySomething(h human) {
	fmt.Println("Hi from the human")
	h.speak()
}
