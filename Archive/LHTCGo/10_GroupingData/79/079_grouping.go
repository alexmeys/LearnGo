package main

import "fmt"

func main() {
	// composite literal
	// a map
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m)
	fmt.Println(m["James"])

	// not stored value will result in 0 value (nil)
	fmt.Println(m["random"])

	// You can check data before doing with the following:
	v, ok := m["random"]
	fmt.Println(v)
	// false
	fmt.Println(ok)

	// more used and condensed:
	if v, ok := m["random"]; ok {
		fmt.Println("Does exist:", v)
	}
}
