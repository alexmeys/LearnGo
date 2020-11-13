// Package test for fun
package test

import (
	"fmt"
	"strings"
)

// global variable
var Name string 

// function who is called to enter name
func Who(){
	fmt.Printf("Enter name: ")
	fmt.Scan(&Name)
}

// function Up is called to go to upper
// this is a fast implementation
func Up(name string) string {
	name = strings.ToUpper(name)
	return name
}

// function Caps is called to go to upper
// this is a slow implementation and missing lower/upper mix
// just to test benchmarks. 
func Caps(name string) string{
	var x string = ""
	for i := 0; i < len(name); i++{
		x += string(name[i]-32)
	}
	return x
}