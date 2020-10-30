package main

import "fmt"

func main() {
	// 33-122 print num : letter
	// you could use %#U in below code for ASCII code + the letter between 'a'
	// however it's nicer to split things up and print it with c.

	for i := 33; i <= 122; i++ {
		fmt.Printf("%d:\t%U:\t%c\n", i, i, i)
	}

	fmt.Println("All Done")
}
