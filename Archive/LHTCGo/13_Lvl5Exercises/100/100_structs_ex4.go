package main

import "fmt"

func main() {

	cd1 := struct {
		name    string
		numbers uint8
		contact map[string]int
		Tracks  []string
	}{
		name:    "Top10",
		numbers: 10,
		contact: map[string]int{
			"Mr Beat": 123,
			"Ms Beat": 456,
		},
		Tracks: []string{"num1", "num2", "num3", "num4", "num5", "num6", "num7", "num8", "num9", "num10"},
	}

	fmt.Println("Name:", cd1.name)
	fmt.Println("Numbers:", cd1.numbers)

	for k, v := range cd1.contact {
		fmt.Printf("\t%v %v\n", k, v)
	}

	for _, v := range cd1.Tracks {
		fmt.Printf("\t\t%v\n", v)
	}
}
