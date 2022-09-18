package main

import "fmt"

func main() {

	s1 := foo()
	fmt.Println(s1)
	s2 := bar()

	fmt.Printf("%T\n", s2)

	x := s2()

	fmt.Println(x)
}

func foo() string {

	s := "Hello World"
	return s
}

func bar() func() int {

	return func() int {
		return 451
	}
}
