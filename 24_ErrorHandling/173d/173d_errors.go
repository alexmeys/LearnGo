package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("nofile.txt")
	if err != nil {
		// exits when failed
		log.Fatalln(err)
	}
}
