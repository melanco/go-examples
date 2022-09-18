package main

import (
	"fmt"
)

func main() {

	james := []string{"James", "Bond", "Shaken, not strirred"}

	penny := []string{"Miss", "Moneypenny", "Helloooo, James"}

	combined := [][]string{james, penny}

	fmt.Println(combined)

	for i, v := range combined {
		for y, u := range v {

			fmt.Println("Dans le slice", i, "et dans l'index", y, "il y a la valeur", u)

		}
	}

	//Autre solution

	for _, comb := range combined {
		for _, val := range comb {
			fmt.Println(val)
		}
	}
}
