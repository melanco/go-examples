package main

import "fmt"

func main() {

	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m["James"])
	fmt.Println(m["Olivier"])

	v, ok := m["Olivier"]

	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Olivier"]; ok {
		fmt.Println("THIS IS THE IF PRINT", v)
	}

}
