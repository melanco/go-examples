package main

import "fmt"

func main() {

	t1 := struct {
		couleur []string
		taille  string
		plain   bool
		motifs  map[string]string
	}{
		couleur: []string{"Noir", "Rouge"},
		taille:  "5XL",
		plain:   false,
		motifs: map[string]string{
			"Zigzag":  "Bleu",
			"Carrée":  "Rouge",
			"Circles": "Green",
		},
	}

	for _, val := range t1.couleur {
		fmt.Println(val)
	}
	for i, v := range t1.motifs {
		fmt.Println(i, v)
	}
}
