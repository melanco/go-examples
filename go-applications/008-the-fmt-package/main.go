package main

import "fmt"

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	fmt.Printf("%b\t%x\t%#X\n", y, y, y)
	s := fmt.Sprintf("%b\t%x\t%#X\n", y, y, y)
	fmt.Printf("%v\n", y)
	fmt.Println(s)
}
