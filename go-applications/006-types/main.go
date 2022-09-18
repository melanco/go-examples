package main

import "fmt"

var y = 42

//DECLARED THAT Z IN OF TYPE STRING.
var z string = "Salut"
var a string = `Oliver à dis 

Ques qui spasse

"Salut"`

func main() {
	//	fmt.Println(y)
	fmt.Printf("%T\n", y)
	//	fmt.Println(z)
	fmt.Printf("%T\n", z)

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	// z = 43
	// fmt.Println(z)
	// fmt.Printf("%T\n", z)
}
