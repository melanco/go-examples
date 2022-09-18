// Create a for loop using this syntax
// - For condition
// Have it print out the years you've been alive.

package main

import "fmt"

func main() {

	x := 1992
	y := 0
	for x <= 2022 {
		x++
		y++
	}
	fmt.Printf("You've been alive for %d years\n", y)
}
