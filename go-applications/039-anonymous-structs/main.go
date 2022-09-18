package main

import "fmt"

func main() {

	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   35,
	}

	fmt.Println(p1)
}
