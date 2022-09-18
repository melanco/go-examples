package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type burgers struct {
	boeuf  bool
	poulet bool
}

type pizza struct {
	crouteFarcie bool
	crouteMince  bool
}

type bouffe struct {
	pizza
	burgers
	liqueur string
}

type order interface {
	menu()
}

func condiment(b bouffe) []string {

	fmt.Println("Quel condiment voulez vous sur votre burger?")

	return
}

func (b bouffe) menu() {

	if b.pizza.crouteFarcie == true {
		fmt.Println("Vous avez commandez une pizza croute farcie!")
	} else {
		fmt.Println("Vous avez commandez une pizza croute mince!")
	}

	if b.burgers.boeuf == true {
		fmt.Println("Vous avez commandez un burger de boeuf!")
	} else {
		fmt.Println("Vous avez commandez un burger de poulet!")
	}

	fmt.Println("Voulez vous choisir vos condiment? (Tapez Oo/Yy pour oui ou Nn pour non")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	// remove the delimeter from the string
	input = strings.TrimSuffix(input, "\n")

	switch input {

	case "Y", "y", "o", "O", "oui", "OUI", "yes", "YES":
		condiment(b)

	default:
		fmt.Println("J'espère que vous aller aimer votre repas sec :)")
	}
}

func main() {

	b1 := bouffe{
		pizza: pizza{
			crouteMince:  false,
			crouteFarcie: true,
		},
		burgers: burgers{
			boeuf:  true,
			poulet: false,
		},
		liqueur: "coke",
	}

	fmt.Println(b1)

	b1.menu()

}
