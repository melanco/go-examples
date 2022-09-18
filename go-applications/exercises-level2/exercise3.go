// TYPED AND UNTYPED CONSTANTS

package main

import "fmt"

const (
	a = 42
	b = "Salut"
)

const g float64 = 67.54

func main() {

	fmt.Println("Salut")

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", g)
}

// Solution :
// const (
// 		a = 42
//		b int = 42
// )
