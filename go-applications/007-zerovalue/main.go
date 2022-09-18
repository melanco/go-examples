package main

import "fmt"

var y string
var z int

func main() {
	// DÉCLARER UNE VARIABLE POUR ÊTRE D'UN CERTAIN TYPE
	// ENSUITE ASSIGNER UNE VALEUR DE CE TYPE À LA VARIABLE EN QUESTION
	fmt.Println("salut", y, "fin de la string")
	fmt.Printf("%T\n", y)

	y = "Bond, james Bond"

	// fmt.Println(y)
	// fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
