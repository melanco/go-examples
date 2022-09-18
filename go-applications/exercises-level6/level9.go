package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Hi, my name is", p.first, "and my last name is", p.last)
}

func main() {

	p1 := person{
		first: "Oli",
		last:  "M",
		age:   29,
	}

	p1.speak()
}
