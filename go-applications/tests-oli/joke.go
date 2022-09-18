package main

import "fmt"

type coffee struct {
	beanType  string
	beanRoast string
}

type coffeePourer struct {
	coffee
	wantCoffee bool
}

type drinker interface {
	drink()
}

func (c coffee) drink() {
	fmt.Println("I want", c.beanType, c.beanRoast, "- for my coffee")
}
func (cP coffeePourer) drink() {
	if cP.wantCoffee == true {
		fmt.Println("Enjoy your coffee!")
	}
}

func main() {
	c1 := coffee{
		beanType:  "Arabica",
		beanRoast: "Dark Roast",
	}

	c1.drink()
}
