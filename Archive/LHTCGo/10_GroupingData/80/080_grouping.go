package main

import "fmt"

func main() {
	m := map[string]int{
		"Alex": 102,
		"R1":   103,
		"R2":   101,
	}

	fmt.Println(m)

	fmt.Println(m["Alex"])
	fmt.Println(m["NoExist"])

	if v, ok := m["NoExist"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("It does not exist in the map...")
	}

	// Add todd
	m["Todd"] = 34
	m["R3"] = 135
	// Print all out
	fmt.Println("")
	for k, v := range m {
		fmt.Println(k, v)
	}

}
