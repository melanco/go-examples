package main

import "fmt"

type circle struct {
	radius float64
	pi     float64
}

type square struct {
	side float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	//On calcule l'air d'un cercle

	squareArea := s.side * s.side
	return squareArea

}

func (c circle) area() float64 {
	//On calcul l'aire d'un cercle

	circleRadiusSquared := c.radius * c.radius
	circleArea := c.pi * circleRadiusSquared

	return circleArea
}

func info(s shape) {

	fmt.Println(s.area())

}

func main() {

	square1 := square{
		side: 15,
	}

	circle1 := circle{
		radius: 7.5,
		pi:     3.1415,
	}

	info(square1)
	info(circle1)

	// x := square1.area(square1)
	// y := circle1.area(circle1)

	// fmt.Println(x)
	// fmt.Println(y)
}
