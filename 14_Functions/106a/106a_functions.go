package main

import "fmt"

type human struct {
	name string
	age  uint8
}

type male struct {
	gender string
	says   []string
	human
}

type female struct {
	gender string
	says   []string
	human
}

type person interface {
	stats()
}

func main() {
	fmt.Println("")

	h1 := male{
		human: human{
			name: "Jefferson",
			age:  20,
		},
		gender: "Male",
		says:   []string{"Hi", "1,2,3"},
	}

	h2 := female{
		human: human{
			name: "Everson",
			age:  21,
		},
		gender: "Female",
		says:   []string{"Ola", "3,2,1"},
	}

	pmMix(h1)
	h1.stats()
	fmt.Println("")
	pmMix(h2)
	h2.stats()
}

func (st female) stats() {
	fmt.Println("Name:", st.name)
	fmt.Println("Age:", st.age)
	fmt.Println("Gender:", st.gender)
	fmt.Println("Says:", st.says)
}
func (st male) stats() {
	fmt.Println("Name:", st.name)
	fmt.Println("Age:", st.age)
	fmt.Println("Gender:", st.gender)
	fmt.Println("Says:", st.says)
}
func pmMix(p person) {
	switch p.(type) {
	case male:
		fmt.Println("Output from polymorphis - interfaces:")
		fmt.Println(p)
		fmt.Println("")
	case female:
		fmt.Println("Output from polymorphis - interfaces:")
		fmt.Println(p)
		fmt.Println("")
	default:
		fmt.Println("No gender defined type")
	}
}
