package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Salut ça va?")
	fmt.Fprintln(os.Stdout, "Salut ça va?")
	io.WriteString(os.Stdout, "Salut ça va?")

}
