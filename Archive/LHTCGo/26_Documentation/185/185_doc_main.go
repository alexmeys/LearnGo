package main

import (
	"fmt"

	"github.com/alexmeys/LearnGo/26_Documentation/185/mymath"
)

func main() {
	fmt.Println("2+3=", mymath.Sum(2, 3))
	fmt.Println("2+3=", mymath.Sum(4, 7))
	fmt.Println("2+3=", mymath.Sum(9, 8))
}
