package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for k := 0; k < 3; k++ {
			fmt.Printf("\t\t The inner loop: %d\n", k)

		}
	}
}
