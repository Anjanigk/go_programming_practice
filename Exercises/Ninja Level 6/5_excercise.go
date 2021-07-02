package main

import (
	"fmt"
	"math"
)

//Method
type circle struct {
	radius float64
}
type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.length * s.length
}

func main() {
	cir := circle{
		10,
	}
	sq := square{
		10,
	}
	fmt.Println(cir.area())
	fmt.Println(sq.area())
}
