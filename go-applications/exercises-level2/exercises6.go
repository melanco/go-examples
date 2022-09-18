package main

import "fmt"

const (
	a = 2018 + iota
	b = 2018 + iota
	c = 2018 + iota
	d = 2018 + iota
)

func main() {

	fmt.Println(a, b, c, d)
}
