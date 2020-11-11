package main

import (
	"fmt"
)

type customErr struct {
	msg string
}

func main() {

	myErr := customErr{
		msg: "You've got an error",
	}

	a := 10
	foo(a, myErr)

}

func (ce customErr) Error() string {
	return fmt.Sprintf("Here is the error: %v", ce.msg)
}

func foo(a int, c error) {
	fmt.Println(a)
	fmt.Println("error:", c)
}
