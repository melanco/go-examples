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

type human interface {
	speak()
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, "- the secretAgent speak")
}
func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "- the person speak")
}

func bar(h human) {

	switch h.(type) { //This is an assert; asserting "h" is of this type
	case person:
		fmt.Println("Je suis une personne!", h.(person).first, h.(person).last)
	case secretAgent:
		fmt.Println("Je suis un agent secret!", h.(secretAgent).first, h.(secretAgent).last)
	}

}

type hotdog int

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

	p1 := person{
		first: "Doctor",
		last:  "Yes",
	}
	fmt.Println(sa1)
	fmt.Println(sa2)
	sa1.speak()
	sa2.speak()
	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)

	// CONVERSION

	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
