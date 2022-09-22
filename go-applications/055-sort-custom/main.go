package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age   int
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

type ByString []person

func (a ByString) Len() int           { return len(a) }
func (a ByString) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByString) Less(i, j int) bool { return a[i].first < a[j].first }

func main() {
	p1 := person{"James", 32}
	p2 := person{"Moneypenny", 27}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}
	p5 := person{"Amen", 82}

	people := []person{p1, p2, p3, p4, p5}

	fmt.Println(people)
	sort.Sort(ByString(people))
	fmt.Println(people)
}
