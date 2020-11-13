package main

import "os"

func main() {
	_, err := os.Open("nofile.txt")
	if err != nil {
		panic(err)
	}
}
