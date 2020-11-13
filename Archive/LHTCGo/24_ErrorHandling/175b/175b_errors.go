package main

import (
	"errors"
	"fmt"
	"log"
)

var errmath = errors.New("Math issue, square root below 0 not possible")

func main() {
	fmt.Printf("%T\n", errmath)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errmath
	}
	return 42, nil
}
