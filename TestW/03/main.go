package main

import "fmt"

type about struct {
	name    string
	version uint8
}

func main() {
	newS := about{
		name:    "MeSoft",
		version: 1,
	}

	newS.printVersion()

}

func (a about) printVersion() {
	fmt.Println(a.name)
	fmt.Println(a.version)
}
