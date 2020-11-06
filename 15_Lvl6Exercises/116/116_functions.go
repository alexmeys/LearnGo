package main

import "fmt"

func main() {
	p1 := person{
		"Jef",
		"Jefferson",
		20,
	}

	p1.speak()
}

type person struct {
	first string
	last  string
	age   uint8
}

func (st person) speak() {
	fmt.Println("Name:", st.first, st.last)
	fmt.Println("Age:", st.age)
}
