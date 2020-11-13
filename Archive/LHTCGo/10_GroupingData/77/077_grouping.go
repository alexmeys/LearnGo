package main

import "fmt"

func main() {

	// Original syntax:
	//x := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}

	// with make we say (TYPE []int (array int), sizeSlice (items of slice), sizeArray (hosting the slices))
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i, v := range x {
		v = (11 * i)
		x[i] = v
	}

	fmt.Println(x)

	// This above code makes it faster and better for the cpu, since we know the size.
	// what happens if we overflow that predefined size?

	// our slice gorws from 10 max to 11, the array below that keeps 12
	x = append(x, 111)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 112)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// Here we go out of space to put items since we defined 12 for max array
	// what happens: the array doubles in size it would print:
	// the full array - 13 (slice items) - 24 (Array size)
	// Here our code would slow down since it needs to re-create the slice
	x = append(x, 113)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
