package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

type human interface {
	speak()
}

func (*person) speak() {
	fmt.Println("I am a human")
}

func main() {
	fmt.Println("Hello, playground")
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	//Ne marche pas à cause des method sets
	// saySomething(p1)
	saySomething(&p1)

}

func saySomething(h human) {

	h.speak()
}
