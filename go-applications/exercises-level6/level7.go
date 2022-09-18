//Create a func which returns a func
// Assign the returned func to a variable
// Call the returned func from that variable

package main

import "fmt"

func main() {

	xi := foo()
	xi()
}

func foo() func() int {

	return func() int {
		x := 451
		fmt.Println(x)
		return x
	}
}
