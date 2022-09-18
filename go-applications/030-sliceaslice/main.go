package main

import "fmt"

func main() {

	x := []int{78, 321, 32, 432543, 3675, 43543, 654, 765}

	// Commence à afficher à partir de l'index 1
	fmt.Println(x[1:])
	// Commence à afficher à l'index 2 et affiche jusqu'à l'index 6 mais sans l'inclure
	fmt.Println(x[2:6])

	// Autre manière d'afficher tout les valeurs mais sans "range"
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
}
