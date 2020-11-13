package main

import "fmt"

func main() {

	xs1 := []string{"James", "Bond", "Shaken not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Hello James"}

	xxs1 := [][]string{xs1, xs2}
	fmt.Println(xxs1)

	for k, v := range xxs1 {
		fmt.Println("Record:", k)
		for j, v2 := range v {
			fmt.Printf("\tIndex pos: %v\t value:%v\n", j, v2)
		}
	}
}
