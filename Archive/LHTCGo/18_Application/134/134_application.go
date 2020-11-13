package main

import (
	"fmt"
	"sort"
)

type person struct {
	First string
	Age   int
}

func main() {
	p1 := person{"James", 32}
	p2 := person{"Moneypenny", 27}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}

	people := []person{p1, p2, p3, p4}

	fmt.Println("")
	fmt.Println("Unsorted:")
	fmt.Println(people)

	fmt.Println("")
	fmt.Println("Sorted by Age:")
	sort.Sort(byage(people))
	fmt.Println(people)

	fmt.Println("")
	fmt.Println("Sorted by Name:")
	sort.Sort(byname(people))
	fmt.Println(people)

}

type byname []person

func (bn byname) Len() int           { return len(bn) }
func (bn byname) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn byname) Less(i, j int) bool { return bn[i].First < bn[j].First }

type byage []person

func (a byage) Len() int           { return len(a) }
func (a byage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byage) Less(i, j int) bool { return a[i].Age < a[j].Age }
