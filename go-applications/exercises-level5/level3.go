package main

import "fmt"

type vehicule struct {
	doors int
	color string
}

type truck struct {
	vehicule
	fourWheels bool
}

type sedan struct {
	vehicule
	luxury bool
}

func main() {

	t1 := truck{
		vehicule: vehicule{
			doors: 2,
			color: "black",
		},
		fourWheels: false,
	}

	s1 := sedan{
		vehicule: vehicule{
			doors: 4,
			color: "red",
		},
		luxury: true,
	}

	fmt.Println(t1, s1)

	fmt.Println(s1.luxury, s1.doors)

}
