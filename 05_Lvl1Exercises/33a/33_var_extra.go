package main

import "fmt"

func main() {
	fmt.Println("Lets get started")

	var voornaam string = "Mr"
	naam := "Go"
	leeftijd := 102

	naam = "NoGo"
	leeftijd++

	s := fmt.Sprintln("Hello", naam, "I see your first name is", voornaam)
	s += fmt.Sprintf("Are you really %v%v\n", leeftijd, "?")
	fmt.Println(s)

	leeftijd--
	var echteLeeftijd float64 = float64(leeftijd)

	fmt.Println("I see you cheating, you are only", (echteLeeftijd/1.3 + 4))

	fmt.Println(fmt.Sprintf("%v", "That will be all"))

}
