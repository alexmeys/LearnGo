package main

import "fmt"

type person struct {
	fname string
	lname string
	age   uint8
}

func main() {
	mySlice := []string{"hi", "there", "how", "are", "you"}
	updateSlice(mySlice)
	fmt.Println(mySlice)

	h1 := person{
		fname: "Jef",
		lname: "Jefferson",
		age:   20,
	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println(h1)
	h1.updateName("John")
	fmt.Println("")
	fmt.Println(h1)
	fmt.Println("")
}

// struct needs a T -> *T or *T -> *T
// in passing T -> *T Go will figure it out, *T -> *T is defined so that works
// *T -> T would not work.
func (p *person) updateName(newName string) {
	p.fname = newName
}

// slice does not need this, because of being a reference type (details ch7 pointers)
func updateSlice(s []string) {
	s[0] = "Bye"
}
