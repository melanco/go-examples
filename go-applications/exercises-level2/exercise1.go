package main

import "fmt"

func main() {
	x := 42
	fmt.Printf("%X\n", x)
	fmt.Printf("%#x\n", x)
	fmt.Printf("%b\n", x)
	// La façon corrigée

	fmt.Printf("%d\t%b\t%#x", x, x, x)

}
