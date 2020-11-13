package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	// Capitals needed otherwise empty struct values as result in marshal
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   28,
	}
	// var to hold all the structs (slice of person struct)
	people := []person{p1, p2}
	// Regular print
	fmt.Println(people)

	// store marshal in bs (byte slice), and catch err if any evaluate
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	// Print the byte code (112 127 ..) in string value (abc)
	fmt.Println(string(bs))

}
