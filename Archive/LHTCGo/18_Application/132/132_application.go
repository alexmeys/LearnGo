package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hi there")
	fmt.Fprintln(os.Stdout, "How are you doing?")
	io.WriteString(os.Stdout, "Bye")
}
