package main

import "fmt"

func main() {

	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	x := foo(xi...)

	y := bar([]int{1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println(x)
	fmt.Println(y)
}

func foo(x ...int) int {

	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum
}

func bar(y []int) int {

	barSum := 0
	for _, j := range y {
		barSum += j
	}

	return barSum
}
