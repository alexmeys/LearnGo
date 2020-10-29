package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Runtime is handy for stuff to check, lets see some examples:")
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
