package main

import "fmt"

func main() {
	foo()

	// anonymous call function

	func() {
		fmt.Println("Anonymous call func")
	}()

	// last one is the () call for the parameters
	// eg 42 here, one argument

	func(x int) {
		fmt.Println("The meaning of life:", x)
	}(42)

}

//identiefied as foo
func foo() {
	fmt.Println("foo ran")
}
