package main

import "fmt"

func main() {

	func() {
		fmt.Println("Ceci est ma fonction anonyme!")
	}()
}