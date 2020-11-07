package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"FName"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	// string of data
	s := `[{"FName":"James",
	"Last":"Bond",
	"Age":32},
	{"Fname":"Miss",
	"Last":"Moneypenny",
	"Age":28}]`

	// convert to byte as it is needed by unmarshal
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	//people := []person{} -- is the same as below, somehow the below is more clear
	var people []person

	// unmarshal use the byte slice to add to the address of people (&) what stores a slice of person
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	// print it all out in one Go
	fmt.Println("All data:", people)
	// range over the items fields like a regular struct
	for i, v := range people {
		fmt.Println("\n--- Person NUM", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

}
