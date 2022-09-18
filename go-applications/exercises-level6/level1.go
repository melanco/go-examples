package main

import "fmt"

type bothThings struct {
	salut  string
	nombre int
}

func main() {

	xi := foo()
	x, s := bar()

	fmt.Println(xi)
	fmt.Println(x, s)
}

// func foo() int {

// 	x := 42

// 	return x
// }

// func bar() bothThings {

// 	x := bothThings{

// 		salut:  "salut",
// 		nombre: 42,
// 	}

// 	return x
// }

// autre solution

func foo() int {
	return 42
}

func bar() (int, string) {
	return 42, "SAlut la!"
}
