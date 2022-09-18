package main

import "fmt"

func main() {

	x := 44

	if x == 42 {
		fmt.Println("42!")
	} else if x == 41 {
		fmt.Println("41!")
	} else if x == 43 {
		fmt.Println("43!")
	} else if x == 44 {
		fmt.Println("44!")
	} else {
		fmt.Println("C'est quoi s'taffaire la")
	}
}
