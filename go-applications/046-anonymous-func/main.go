package main

import "fmt"

func main() {

	foo()
	// fmt.Println("Validations")

	func(x int) {
		fmt.Println(x)
	}(42)
}

func foo() {

	fmt.Println("SAlut")
}
