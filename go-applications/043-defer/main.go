package main

import "fmt"

func main() {
	defer foo()
	bar()
	tang()
}

func foo() {

	fmt.Println("foo")

}

func bar() {

	fmt.Println("bar")

}

func tang() {

	fmt.Println("tang")

}
