package main

import "fmt"

func main() {
	fmt.Println("")

	//test := make([]string, 4, 5)

	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)

	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	// Multi Dimensional slices
	md := [][]string{jb, mp}
	fmt.Println(md)

	// using data from MD slice
	fmt.Println(md[0][1])
	fmt.Println(md[1][1])

}
