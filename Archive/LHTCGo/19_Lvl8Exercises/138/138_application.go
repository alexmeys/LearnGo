package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     uint8
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken not stirred",
			"Youth is no guarantee of innovation",
			"In this majesty's royal service",
		},
	}
	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James.",
			"I would really prefer to be secret agent myself.",
		},
	}
	u3 := user{
		First: "M",
		Last:  "Hmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	// fmt.Println(users)
	// err := json.NewEncoder(os.stdout).encode(users) also works, but this is split up

	jsonw := json.NewEncoder(os.Stdout)
	err := jsonw.Encode(users)
	if err != nil {
		fmt.Println(err)
	}

}
