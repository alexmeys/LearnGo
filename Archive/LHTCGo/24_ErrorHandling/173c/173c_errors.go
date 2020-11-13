package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("Log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	// sets the output log and writes it to f
	log.SetOutput(f)

	// open not existing file:
	f2, err := os.Open("nofile.txt")
	if err != nil {
		log.Println("Error:", err)
	}
	defer f2.Close()
	fmt.Println("Check the log.txt file in the dir")
}
