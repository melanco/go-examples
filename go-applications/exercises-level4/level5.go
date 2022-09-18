package main

import "fmt"

func main() {

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// To delete form a slice we use Append along with slicing.

	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}
