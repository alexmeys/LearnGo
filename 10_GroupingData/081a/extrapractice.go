package main

import "fmt"

func main(){
	// Using Map
	fmt.Println("")
	m := map[string]int{
		"Alex" : 0,
		"Ben" : 1,
		"Cedric": 2,
	}

	fmt.Println("Map: ", m)
	fmt.Println("Adding Jef to map")
	m["Jef"] = 3

	fmt.Println("New Map: ", m)

	fmt.Println("Remove Ben")
	delete(m, "Ben")

	fmt.Println("New Map: ", m)

	fmt.Println("")

	// Using slices
	xi := []int{2,4,6,8,9,10}
	fmt.Println(xi)

	// Append 12 to the end
	xi = append(xi, 12)

	fmt.Println(xi)

	// remove 9
	xi = append(xi[0:4], xi[5:]...)
	fmt.Println(xi)


	fmt.Println("")

	// faster slices
	yi := make([]int, 4, 5)

	fmt.Println("Length slice:",len(yi))
	fmt.Println("Capacity array:", cap(yi))

	for i, v := range yi{
		if i %2 == 0{
			v = 0
			yi[i] = v
		} else {
			v = 1
			yi[i] = v
		}
	}

	fmt.Println(yi)

}
