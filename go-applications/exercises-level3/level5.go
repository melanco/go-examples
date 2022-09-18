// Print out the reminder (modulo) which is found for each number between 10 and 100 when it is devided by 4

package main

import "fmt"

func main() {

	for i := 10; i <= 100; i++ {
		y := i % 4
		fmt.Printf("%d\t %d\n", i, y)
	}

	// Autre solution

	for y := 10; y <= 100; y++ {
		fmt.Printf("Quand on prend le nombre %v, et qu'on fait modulo 4 dessus, ça donne %v\t\n", y, y%4)
	}
}
