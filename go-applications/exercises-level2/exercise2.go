package main

import "fmt"

func main() {

	nb1 := 43
	nb2 := 45

	a := (nb1 <= nb2)
	b := (nb1 >= nb2)
	c := (nb1 != nb2)
	d := (nb1 < nb2)
	e := (nb1 > nb2)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

}
