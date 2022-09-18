package main

import "fmt"

func main() {

	x := []int{4, 5, 6, 7, 8}

	// avoir le lenght de l'array
	fmt.Println(len(x))
	//	fmt.Println(cap(x))
	fmt.Println(x)
	// Accéder à des choses directement sur un index
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])

	// pour looper sur le slice
	for i, v := range x {
		fmt.Println(i, v)
	}
}
