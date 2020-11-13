package main

import "fmt"

func main() {
	var myArr [5]int
	// static typed input data
	// faster woud be myArr := [5]int{10, 11, 12,...}
	myArr[0] = 10
	myArr[1] = 11
	myArr[2] = 12
	myArr[3] = 13
	myArr[4] = 14

	for k, v := range myArr {
		fmt.Println("Key: ", k, "Value: ", v)
	}

	fmt.Printf("Type of Array: %T\n", myArr)
}
