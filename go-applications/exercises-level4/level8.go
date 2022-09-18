package main

import "fmt"

func main() {

	//`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
	//`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
	//`no_dr`, `Being evil`, `Ice cream`, `Sunsets`

	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	for i, v := range m {
		fmt.Println("Personne courrante :", i)
		for j, val := range v {
			fmt.Printf("\tÀ l'index: %v \t La valeur est %v\n", j, val)
		}
	}
}
