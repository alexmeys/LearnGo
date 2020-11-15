package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// removing the variable in receiver, as it is empty...
func (englishBot) getGreeting() string {
	// imagine specific custom english code here
	return "Hello there!"
}

func (spanishBot) getGreeting() string {
	// imagine specific custom spanish code here
	return "Hola!"
}
