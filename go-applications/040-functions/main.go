package main

import "fmt"

func main() {
	foo()
	bar("Olivier")
	s1 := woo("Miss Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Flemming")

	fmt.Println(x, y)
}

func foo() {

	fmt.Println("Hello from foo")
}

func bar(s string) {
	fmt.Printf("Salut mon cher %v\n", s)
}

func woo(m string) string {

	return fmt.Sprintln("Hello from woo,", m)
}

func mouse(fn, ln string) (string, bool) {
	a := fmt.Sprintln(fn, ln, `, says "hello"`)
	b := true

	return a, b
}
