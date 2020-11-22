package main

import (
	"LearnGo/GoCBCMP/02_Hello/09_excs/car"
	"fmt"
)
func main(){
	stateEngine := car.Engine()
	fmt.Println(stateEngine)
}
