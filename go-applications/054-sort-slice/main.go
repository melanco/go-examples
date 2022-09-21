package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Salut ! ")

	xi := []int{1, 4, 3, 5, 7, 9, 2, 6}
	xs := []string{"Café", "Beigne", "Muffin", "Espresso", "Vodka"}

	// on print le avant
	fmt.Println(xi)
	sort.Ints(xi)
	//on print le après
	fmt.Println(xi)

	fmt.Println("---------")

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
