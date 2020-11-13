package main

import "fmt"

func main() {
	x := []int{99, 88, 77, 66}
	fmt.Println(x)

	// append
	x = append(x, 55, 44, 33, 22, 11)
	fmt.Println(x)

	// append from another slice Variadic parameter
	// y... = use all the values of y and put them behind x appended
	// ... is needed because it would not compile and complain about not being an int type "y is not of type int"
	y := []int{3513, 6354, 5635, 2213}
	x = append(x, y...)

	fmt.Println(x)

}
