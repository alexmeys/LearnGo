package main

import "fmt"

func main() {
	fmt.Println("Starting my GO jouney!")

	foo()

	fmt.Println("After foo()")

	for i := 0; i < 50; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("...Foo is staring up...", 42, true)

}

func bar() {
	n, err := fmt.Println("The end...")
	fmt.Println("Number of bytes:", n)
	fmt.Println("Errors encountered:", err)
}
