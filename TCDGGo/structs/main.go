package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Jimmerson",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 3400,
		},
	}

	jim.updateName("Jimmy")
	fmt.Println("")
	jim.print()
	fmt.Println("")

	alex.updateName("Alexander")
	fmt.Println("")
	alex.print()
	fmt.Println("")

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (pointerToPerson *person) print() {
	fmt.Printf("%+v", pointerToPerson)
}
