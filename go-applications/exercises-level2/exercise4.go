package main

import "fmt"

func main() {
	a := 42
	fmt.Printf("%d\t%b\t%#x\t\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\t%b\t%#x\t", b, b, b)
}
