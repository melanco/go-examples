//How do you test this? It is good to separate your "domain" code from the outside world (side-effects).
//The fmt.Println is a side effect (printing to stdout) and the string we send in is our domain.

// Ici on créer une fonction qui retourne un string, et on valide dans hello_test.go que cette string est bonne.
// Il faut absolument rouler un go mod init pour que le module comprenne notre environnement.
// rouler go test dans le terminal pour voir le résultat.

package main

import "fmt"

func Hello() string {
	return "Hello, World!"
}

func main() {

	fmt.Println(Hello())

}
