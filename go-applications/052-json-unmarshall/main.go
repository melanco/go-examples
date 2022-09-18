package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {

	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Monneypenny","Age":27}]`
	// On converti le json en une slice de byte
	bs := []byte(s)

	// On peut créer une slice de person
	//people := []person{}
	//Ou explicitement créer le type people qui est une slice de person
	var people []person

	err := json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("all the data", people)

	for i, v := range people {
		fmt.Println("\n---- Person number", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

	// fmt.Printf("%T\n", s)
	// fmt.Printf("%T\n", bs)

}
