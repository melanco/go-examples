package main

import "fmt"

type person struct {
	first   string
	last    string
	age     int
	address string
}

func changeMe(p *person) (s string) {

	p.address = "420 de bourges"

	fmt.Println(&p.address)
	return p.address
}

func main() {

	p1 := person{
		first:   "Oli",
		last:    "M",
		age:     29,
		address: "1551 donnacona",
	}

	fmt.Println(&p1.address)
	fmt.Println(p1.address)

	modifiedAddress := changeMe(&p1)

	fmt.Println(modifiedAddress)
}
