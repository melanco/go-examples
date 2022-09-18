package main

import "fmt"

func main() {
	fmt.Println("Salut")
	foo()
	fmt.Println("SAlut encore ecnore?")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("SAlut encore")
}

func bar() {
	fmt.Println("Et finalement on sort..")
}
