package main

import (
	"fmt"
	"sort"
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

	fmt.Println("Sort by Age:")
	sort.Sort(byage(users))
	for _, v := range users {
		fmt.Println(v.First)
		fmt.Println(v.Last)
		fmt.Println(v.Age)
		fmt.Println("Sayings:")
		sort.Strings(v.Sayings)
		for _, v2 := range v.Sayings {
			fmt.Println(v2)
		}
		fmt.Println("")
	}

	fmt.Println("")
	fmt.Println("Sort by Last name:")
	sort.Sort(bylast(users))
	for _, v := range users {
		fmt.Println(v.First)
		fmt.Println(v.Last)
		fmt.Println(v.Age)
		fmt.Println("Sayings:")
		sort.Strings(v.Sayings)
		for _, v2 := range v.Sayings {
			fmt.Println(v2)
		}
		fmt.Println("")
	}

}

func (u user) String() string {
	return fmt.Sprintf("\tName:%s %s, Age:%d, %v", u.First, u.Last, u.Age, u.Sayings)
}

type byage []user

func (a byage) Len() int           { return len(a) }
func (a byage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byage) Less(i, j int) bool { return a[i].Age < a[j].Age }

type bylast []user

func (a bylast) Len() int           { return len(a) }
func (a bylast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a bylast) Less(i, j int) bool { return a[i].Last < a[j].Last }
