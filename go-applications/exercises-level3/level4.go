//Create a for loop using this syntax
// for {}
//Have it print out the years you've been alive.

package main

import "fmt"

func main() {

	x := 1992
	y := 2022

	for {
		if x <= y {
			fmt.Println(x)
			x++
		}
	}

}
