package main

import "fmt"

// func (r receiver) identifier(parameters) (returns(s)) { ... }

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s person) speak() {

	fmt.Println("I am", s.first, s.last)

}

func main() {

	sa1 := secretAgent{
		person: person{
			first: "james",
			last:  "bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "miss",
			last:  "moneypenny",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	fmt.Println(sa2)
	sa1.speak()
	sa2.speak()
}
