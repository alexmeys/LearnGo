package main

import "fmt"

func main() {
	m := map[string][]string{
		"Bond_James":     []string{"Shaken not stirred", "Martinis", "Woman"},
		"Moneypenny_Mis": []string{"James bond", "Literature", "Computer Science"},
		"no_dr":          []string{"Being evil", "Ice cream", "Sunsets"},
	}

	m["Alex"] = []string{"Go", "MoreGo"}

	for k, v := range m {
		fmt.Println("Key:", k)
		for _, v2 := range v {
			fmt.Println("\tValue:", v2)
		}
	}

}
