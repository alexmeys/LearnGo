package main

import "fmt"

func main() {
	foo()
	bar("James")

	s1 := woo("Moneypenny")
	fmt.Println(s1)

	x, y := zoo("Ian", "Flemming")
	fmt.Println(x)
	fmt.Println(y)
}

func foo() {
	fmt.Println("Hello from go")
}

// pass by value
func bar(s string) {
	fmt.Println("Hello,", s)
}

// return value to main func woo
func woo(s string) string {
	return fmt.Sprint("Hello from woo,", s)
}

// multiple return values
func zoo(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, " says 'hello'")
	b := false
	return a, b
}
