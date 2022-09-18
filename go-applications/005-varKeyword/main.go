package main

import "fmt"

var y = 44

// DECLARE THAT THERE IS A VARIABLE "z"
// AND THAT THE VARIABLE IS TYPE INTEGER
// ASSIGNS THE 0 VALUE of TYPE INT to "z"
var z int

func main() {
	// short declaration operator
	// DECLARE A VARIABLE AND ASSIGN A VALUE (OF A CERTAIN TYPE)
	x := 42
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	foo()
}

func foo() {
	fmt.Println("again", y)
}
