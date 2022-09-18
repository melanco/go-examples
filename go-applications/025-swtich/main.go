package main

import (
	"fmt"
)

func main() {

	n := "Bond"

	switch n {
	case "Mr. Requin", "Mr. Méchant", "Goldenfinger":
		fmt.Println("Mr. Requin!, Mr Méchant!, Goldenfinger!")
	case "James", "James Bond", "Mr. Bond", "Bond":
		fmt.Println("James, James Bond")
	default:
		fmt.Println("Le cas par défaut si aucun autre case ne passe")
	}
}
