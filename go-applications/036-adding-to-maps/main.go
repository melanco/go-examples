package main

import "fmt"

func main() {

	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}

	fmt.Println(m["James"])
	fmt.Println(m["Ian Flemming"])

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("Value :", v)
		delete(m, "Miss Moneypenny")
	}

	fmt.Println(m)
}
