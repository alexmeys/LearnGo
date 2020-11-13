package main

import "fmt"

func main() {
	type vehicle struct {
		doors uint8
		color string
	}
	type truck struct {
		vehicle
		fourWheel bool
	}
	type sedan struct {
		vehicle
		luxury bool
	}

	v1 := truck{
		vehicle: vehicle{
			doors: 5,
			color: "black",
		},
		fourWheel: true,
	}

	v2 := sedan{
		vehicle: vehicle{
			doors: 3,
			color: "blue",
		},
		luxury: true,
	}

	fmt.Println(v1, v2)

	fmt.Println(v2.doors, v2.luxury)
	fmt.Println(v1.doors, v1.fourWheel)

}
