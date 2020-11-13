package main

import (
	"fmt"

	"github.com/alexmeys/LearnGo/29_Lvl13Exercises/195/p195"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  p195.Years(10),
	}
	fmt.Println(fido)
}
