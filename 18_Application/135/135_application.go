package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Println(bs)

	loginpwd := `password123d`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginpwd))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("All good")

}
