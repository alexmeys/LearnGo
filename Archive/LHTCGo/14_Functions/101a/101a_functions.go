package main

import "fmt"

func main() {
	fmt.Println("Functions")
	abc()
	def(2020)
	s := efg("true")
	fmt.Println(s)

	t, u := hij()
	fmt.Println(t, u)

}

func abc() {
	fmt.Println("Welcome to functions")
}
func def(d int) {
	fmt.Println("The year is:", d)
}
func efg(s string) bool {
	var result bool
	if s == "true" {
		result = true
	} else {
		result = false
	}
	return result
}
func hij() (string, string) {
	return "The", "End"
}
