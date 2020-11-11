package main

import (
	"fmt"
)

func main() {
	n := CalcDogYears()
	fmt.Println("- Dog year converter -")
	fmt.Println("Your dog is:", n*7)
}
