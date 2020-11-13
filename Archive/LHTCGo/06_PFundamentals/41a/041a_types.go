package main

import "fmt"

const (
	// start the first iota = 0 and throw it out
	// from kb -> mb =* 1024 or equals to 10 places shifted over in binary.
	// iota (1) times 10 and shift the 1 10 places
	// iota (2) times 20 and shift the 1 20 places
	// iota (3) times 30 and shift the 1 30 places

	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

}
