package main

import "fmt"

func main() {

	defer foo()
	bar()
	tang()

}

func foo() {
	fmt.Println("Salut")
}

func bar() {

	fmt.Println("Mon beau")
}

func tang() {
	fmt.Println("Ques tu caliss")
}
