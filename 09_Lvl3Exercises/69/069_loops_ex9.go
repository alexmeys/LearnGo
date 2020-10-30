package main

import "fmt"

func main() {
	var favSport = "Skiing"

	switch favSport {
	default:
		fmt.Println("In default, we did not find skiing")
	case "Skiing":
		fmt.Println("Ey! We found skiing")
	}
}
