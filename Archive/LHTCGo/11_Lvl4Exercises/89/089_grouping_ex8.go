package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken not stirred", "Martinis", "Woman"},
		"moneypenny_miss": []string{"James bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	for k, v := range m {
		fmt.Println("Key:", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
	fmt.Println(m)
}
