package main

import "fmt"

type square struct {
	length float32
	width  float32
}
type circle struct {
	pi float32
	r  float32
}

type shape interface {
	area() float32
}

func main() {
	sq1 := square{
		length: 17,
		width:  18,
	}

	cr1 := circle{
		pi: 3.14159,
		r:  17,
	}

	info(sq1)
	info(cr1)

}

func (sq square) area() float32 {
	return (sq.length * sq.width)
}

func (cr circle) area() float32 {
	return (cr.pi * (cr.r * cr.r))
}

func info(s shape) {
	switch s.(type) {
	case circle:
		fmt.Println("The circle area:", s.area())
	case square:
		fmt.Println("The square area:", s.area())
	default:
		fmt.Println("Did not find the object")
	}
}
