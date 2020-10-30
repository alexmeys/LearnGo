package main

import "fmt"

func main() {
	x := 100
	if x == 40 {
		fmt.Println("Our val is 40")
	} else if x > 1 && x <= 100 {
		for i := 0; i <= 100; i++ {
			if i == x {
				fmt.Printf("Found your number it's %d\n", x)
			} else {
				continue
			}
		}
	} else {
		fmt.Println("Your nummber is bigger then 100 or 0")
	}
	fmt.Println("Done")
}
