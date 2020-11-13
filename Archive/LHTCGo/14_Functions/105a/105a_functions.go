package main

import "fmt"

type male struct {
	gender   string
	sentence []string
	human
}

type female struct {
	gender   string
	sentence []string
	human
}

type human struct {
	weight   uint8
	length   uint8
	shoesize uint8
	age      uint8
}

func main() {
	fmt.Println("")
	h1 := female{
		human: human{
			weight:   50,
			length:   160,
			shoesize: 38,
		},
		gender:   "Female",
		sentence: []string{"Hi there", "Says how are you", "Bye bye"},
	}

	h2 := male{
		human: human{
			weight:   70,
			length:   170,
			shoesize: 41,
		},
		gender:   "Male",
		sentence: []string{"Says Hi", "Says Bye"},
	}

	defer h2.maletalk()
	fmt.Println("")
	h1.femaletalk()
	fmt.Println("")

}

func (st male) maletalk() {
	fmt.Println("Gender:", st.gender)
	fmt.Println("Weight:", st.weight)
	fmt.Println("Height:", st.length)
	fmt.Println("Shoesize:", st.shoesize)
	fmt.Println("Famous quotes:")
	for _, v := range st.sentence {
		fmt.Println(v)
	}

}

func (st female) femaletalk() {
	fmt.Println("Gender:", st.gender)
	fmt.Println("Weight:", st.weight)
	fmt.Println("Height:", st.length)
	fmt.Println("Shoesize:", st.shoesize)
	fmt.Println("Famous quotes:")
	for _, v := range st.sentence {
		fmt.Println(v)
	}
}
