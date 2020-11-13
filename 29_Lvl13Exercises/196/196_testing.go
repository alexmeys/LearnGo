package main

import (
	"fmt"

	"github.com/alexmeys/LearnGo/29_Lvl13Exercises/196/p196"
	"github.com/alexmeys/LearnGo/29_Lvl13Exercises/196/p196quote"
)

func main() {
	fmt.Println(p196.Count(p196quote.SunAlso))

	for k, v := range p196.UseCount(p196quote.SunAlso) {
		fmt.Println(v, k)
	}
}
