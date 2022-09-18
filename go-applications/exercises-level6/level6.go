package main

import "fmt"

func main() {

	f := func() {
		fmt.Println("J'assigne une fonction à f")
	}

	fmt.Printf("%T\n", f)

	f()

	// On peut aussi faire quelque chose du genre

	g := f

	fmt.Printf("this is g %T\n", g)
	g()

}
