package main

import "fmt"

type oli int

var x oli

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x := 42
	fmt.Println(x)
}
