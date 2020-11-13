package main

import "fmt"

func main() {
	a := 42
	// Regular - value T prints int
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	// accessing the Pointer (*<type> with &)
	// value T prints *int, type. Point to address where int is stored
	fmt.Println(&a)
	fmt.Printf("%T\n", &a)

	// b = pointer type int of the address a
	var b *int = &a
	fmt.Println(b)

	// *b points to the data stored in address b (which is a reference to a)
	fmt.Println(*b)
	// print the value in one line
	fmt.Println(*&a)
	// change the pointer value b to 43, which referes to a so we print a and get 43
	*b = 43
	fmt.Println(a)

}
