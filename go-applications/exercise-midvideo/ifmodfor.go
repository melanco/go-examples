package main

import "fmt"

func main() {

	// x := 10
	// y := 15
	for i := 0; i <= 100; i++ {

		if i%2 == 0 {
			fmt.Printf("Ce nombre %d est pair!\n", i)
		} else {
			fmt.Printf("Ce nombre %d est impair!\n", i)
		}
	}
}
