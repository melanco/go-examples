package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	favIceCream []string
}

func main() {

	m := map[string]person{
		"Melancon": {
			firstName:   "Olivier",
			lastName:    "Melancon",
			favIceCream: []string{"Fraise", "Vanille"},
		},
		"Gosselin": {
			firstName:   "Annie",
			lastName:    "Gosselin",
			favIceCream: []string{"Chocolat", "Banane"},
		},
	}

	for i, v := range m {
		fmt.Println(i, v)
	}

	// La solution :

	// m := map[string]person{
	// 		p1.lastName: p1,
	//		p2.lastName: p2,
	// {
}
