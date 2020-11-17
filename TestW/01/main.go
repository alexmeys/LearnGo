package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	translate() string
}

func main() {
	fmt.Println("Testscope")
	eb := englishBot{}
	sb := spanishBot{}

	tprint(eb)
	tprint(sb)

	eb.dprint()

}

func (spanishBot) translate() string {
	return "Hola!"
}

func (englishBot) translate() string {
	return "Hello there"
}

func tprint(b bot) {
	fmt.Printf("%v\n", b.translate())
}

func (englishBot) dprint() {
	fmt.Println("This is the string Answer")
}
