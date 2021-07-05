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

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("The area is", s.area())
}
func main() {
	cir := circle{
		10,
	}
	sq := square{
		10,
	}
	info(cir)
	info(sq)
	
}
