package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {

	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}

	p2 := person{
		First: "Miss",
		Last:  "Monneypenny",
		Age:   27,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	// On prend la slice de "person" et on la "Marshal"

	b, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	os.Stdout.Write(b)
}
