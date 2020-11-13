package main

import "fmt"

func main() {
	x := 42
	y := 42.34545
	fmt.Println("Value x:", x)
	fmt.Println("Value y:", y)

	fmt.Printf("Type of x: %T\n", x)
	fmt.Printf("Type of y: %T\n", y)

	// Above is rather wasteful
	// rune = int32
	// byte = uint8
	// overflow protection build-in, in case b for example = 256 (get warning)
	var a rune = 32
	var b byte = 255
	var c uint16 = 16

	fmt.Println("Value a:", a)
	fmt.Println("Value b:", b)
	fmt.Println("Value c:", c)

	fmt.Printf("Type of a: %T\n", a)
	fmt.Printf("Type of b: %T\n", b)
	fmt.Printf("Type of c: %T\n", c)

}
