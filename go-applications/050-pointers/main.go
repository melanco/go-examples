package main

import (
	"fmt"
)

func main() {

	a := 42

	fmt.Println(&a)
	b := &a //gives you the address in memory

	fmt.Println(*b) // gives you the value stored at an address, when you have an address.

	fmt.Printf("%T\n", *b)

	fmt.Println(*&a)

	*b = 43

	fmt.Println(a)
	fmt.Println(&a)
}
