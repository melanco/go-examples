package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	s := `password123`
	s1 := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Println(bs)

	err = bcrypt.CompareHashAndPassword(bs, []byte(s1))

	if err != nil {
		fmt.Println("YOU CAN'T LOGIN!!")
		return
	}

	fmt.Println("YOU'RE LOG'ED IN")

}
