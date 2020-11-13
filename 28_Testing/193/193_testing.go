package main

import (
	"fmt"
	"strings"

	"github.com/alexmeys/LearnGo/28_Testing/193/p193"
)

const s = "We ask ourselves, Who am I to be brilliant, gorgeous talented, faboulous? Actually, who are you not to be? Your playing small does not serve the world. There nothing enlightened about shrinking so that other people won't feel insecure around you. We are all meant to shine, as us; it's in everyone. And as we let our own light shine, we unconsciously give other people permission to do the same. As we are liberated from our own fear, our presence automatically liberates others. - Marianne Williamson"

func main() {
	xs := strings.Split(s, " ")

	for _, v := range xs {
		fmt.Println(v)
	}

	fmt.Printf("\n%s\n", p193.Cat(xs))
	fmt.Printf("\n%s\n\n", p193.Join(xs))

}
