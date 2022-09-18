package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	favIceCream []string
}

func main() {

	p1 := person{
		firstName:   "Olivier",
		lastName:    "Melancon",
		favIceCream: []string{"Fraise", "Vanille"},
	}
	p2 := person{
		firstName:   "Annie",
		lastName:    "Gosselin",
		favIceCream: []string{"Chocolat", "Banane"},
	}

	for i, v := range p1.favIceCream {
		fmt.Println(i, v)
	}

	for y, val := range p2.favIceCream {
		fmt.Println(y, val)
	}

}
