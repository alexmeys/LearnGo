package main

import "fmt"

func main() {

	s := "Alex"

	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// Slice of bytes
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n\n", bs)

	// l letter to utf8
	fmt.Println("UTF8: ")
	fmt.Printf("%#U \n\n", s[1])

	// l letter to hex
	fmt.Println("Hex: ")
	fmt.Printf("%#x \n\n", s[1])

}
