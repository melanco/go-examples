package main

import "fmt"

func main() {

	x := []int{4, 5, 6, 7, 8, 9}

	fmt.Println(x)

	//Ensuite pour append la fonction retourne un slice du même type.

	y := []int{10, 11, 12, 13, 14}

	x = append(x, y...)
	fmt.Println(x)
}
