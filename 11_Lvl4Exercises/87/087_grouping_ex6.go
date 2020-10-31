package main

import "fmt"

func main() {
	sStates := make([]string, 6, 6)

	fmt.Println(sStates)
	fmt.Println(len(sStates))
	fmt.Println(cap(sStates))

	sStates = []string{"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado"}

	for i := 0; i < len(sStates); i++ {
		fmt.Println("Index: ", i, "State:", sStates[i])
	}
}
