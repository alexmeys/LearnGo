package main

import "fmt"

func main() {
	fmt.Println("")

	m := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 6,
		"F": 6,
		"G": 7,
	}

	fmt.Println(m)

	delete(m, "E")
	fmt.Println(m)

	// Add E to correct number
	// will auto sort output on Letter (based on key number)
	m["E"] = 5

	// Deletes unknown key, without error
	delete(m, "H")
	fmt.Println(m)

	if v, ok := m["H"]; ok {
		fmt.Println("Value: ", v)
		delete(m, "H")
	} else {
		fmt.Println("No value for this letter found (replace hardcoded H with var an reference here)")
	}
}
