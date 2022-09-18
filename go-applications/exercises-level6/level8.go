package main

import "fmt"

func main() {

	x := bar("what")
	fmt.Println(x)

	xi := foo(bar)
	fmt.Println(xi)

}

func bar(s string) string {
	x := "the"
	new := s + x + "fuck"
	return new
}

func foo(xi func(s string) string) string {

	t := xi("You wot mate")

	tt := t + "VOILA"
	return tt
}
