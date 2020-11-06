package main

import "fmt"

func main() {
	x := 20
	foo(&x)
	bar(x)
	shiftit(&x)
	var pointsToInt *int = &x
	fmt.Println("Create pointer:", pointsToInt)

	fmt.Printf("Printed hex: %#x\n", pointsToInt)

	var pi float32 = 3.14159
	var radius float32 = 20
	var area float32 = 0

	calcarea(&pi, &radius, &area)
	fmt.Println(area)

}
func foo(x *int) {
	fmt.Println("Address of X:", x)
	fmt.Println("Data of X:", *x)
}

func bar(x int) {
	fmt.Println("Address of X:", &x)
	fmt.Println("Data of X:", x)
}

func shiftit(x *int) {
	fmt.Println("Shifted X by 1 Data:", *x<<1)
	fmt.Println("Shifted X by 1 Address:", x)

}
func calcarea(p *float32, r *float32, a *float32) {
	*a = (*p * *r * *r)
}
