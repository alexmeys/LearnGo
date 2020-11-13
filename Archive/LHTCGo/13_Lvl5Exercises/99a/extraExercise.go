package main

import "fmt"

func main(){
	fmt.Println("Welcome to Go")
	fmt.Printf("%s\n", "lets explore and test")
	
	var a = 123
	//b := "123"

	fmt.Printf("UTF: %#U\n", a)
	fmt.Printf("Dec: %d\n", a)
	fmt.Printf("hex: %#x\n", a)
	//fmt.Printf("string: %s\n", a)
	fmt.Printf("char: %c\n", a)

	c := true
	fmt.Println(c)

	const (
		z = iota
		y
		x
		w
	)
	fmt.Println(x)

	fmt.Println(x<<1)

	if x<<1 == 4 {
		fmt.Println("Jup same")
	} else {
		fmt.Println("Not same")
	}

	var myArr [5]int
	myArr[3] = 12

	fmt.Println("My array", myArr)

	var sl1 = make([]int, 4, 5)

	fmt.Println(len(sl1))
	fmt.Println(cap(sl1))

	for i := 0; i<len(sl1); i++{
		sl1[i] = i+1
	}

	fmt.Println("My slice", sl1)

	m := map[uint8][]string{
		1: []string{"ABC", "DEF","GHI"},
		2: []string{"JKL", "MNO", "PQR"},
		3: []string{"STU", "VWX", "YZ"},
	}

	fmt.Println(m)

	if _, ok := m[0]; ok{
		fmt.Print("Title does exist")
	} else {
		m[0]=[]string{"The", "Alpha", "Beth"}
	}

	fmt.Println(m)

	type human struct{
		fname string
		lname string
		age uint8
		comment []string
		credit map[string][]uint8
	}

	type foreign struct{
		human
		country string
	}

	p1 := foreign {
		human: human {
			fname : "Jef",
			lname : "Jefferson",
			age : 12,
			comment: []string{"Likes Cola", "Likes fries"},
			credit: map[string][]uint8{
				"Won": []uint8{1,2,1,3,1},
				"Losses": []uint8{0},
			},
		},
		country: "Belgium",
	}

	p2 := human {
		fname : "John",
		lname : "Johnson",
		age : 40,
		comment: []string{"Likes Singing", "Lieks writing songs", "dislikes something"},
		credit: map[string][]uint8{
			"Won": []uint8{1,2,3,8,9,1},
			"Losses": []uint8{3,2,5},
		},
	}

	fmt.Println("\n\nWe have:")
	fmt.Println(p1.fname, p1.lname)
	fmt.Println("Age:", p1.age)
	fmt.Println("Likes and dislikes:")
	for _, v := range p1.comment {
		fmt.Printf("\t%v\n", v)
	}

	fmt.Println("Game results:")
	for k, v := range p1.credit {
		for _, v2 := range v {
			fmt.Printf("%v: %v\t", k, v2)
		}
	}

	fmt.Println("\n\nWe have:")
	fmt.Println(p2.fname, p2.lname)
	fmt.Println("Age:", p2.age)
	fmt.Println("Likes and dislikes:")
	for _, v := range p2.comment {
		fmt.Printf("\t%v\n", v)
	}

	fmt.Println("Game results:")
	for k, v := range p2.credit {
		for _, v2 := range v{
			fmt.Printf("%v: %v\t", k, v2)
		}
	}


	fmt.Println("\n\n\nThat's all for today Folks!\n\n")

}
