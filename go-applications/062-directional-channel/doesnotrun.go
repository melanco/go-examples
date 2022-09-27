package main

import "fmt"

func main() {
	c := make(chan<- int, 2)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Printf("------\n")

	fmt.Printf("%T\n", c)
}
