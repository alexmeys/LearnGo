package main

import (
	"fmt"
	"github.com/alexmeys/LearnGo/LHTCGo_Extra/01/test"
)


func main(){
	test.Who()
	fmt.Printf("%s %s%s", "Hello Dear", test.Name,".\n")
	a := test.Caps(test.Name)
	b := test.Up(test.Name)
	
	fmt.Println(a,b)
}

