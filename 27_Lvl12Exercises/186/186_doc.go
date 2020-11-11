package main

import (
	"fmt"

	"github.com/alexmeys/LearnGo/27_Lvl12Exercises/186/calcdogyears"
)

func main() {
	var v int
	fmt.Printf("Enter the age of your dog: ")
	_, err := fmt.Scan(&v)
	if err != nil {
		fmt.Print("Error encountered")
	}
	n := calcdogyears.CalcDogYears(v)
	fmt.Println("- Dog year converter -")
	fmt.Println("Your dog is:", n)
}
