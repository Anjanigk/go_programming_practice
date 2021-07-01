package main

import "fmt"

type person struct {
	fname   string
	lname   string
	flavour []string
}

func main() {
	p1 := person{
		fname:   "John",
		lname:   "Smith",
		flavour: []string{"Vanila", "Choclate", "Butterscoch"},
	}
	p2 := person{
		fname:   "Harry",
		lname:   "Potter",
		flavour: []string{"Strawberry", "Choclate", "Mocha"},
	}
	fmt.Println(p1.fname)
	fmt.Println(p1.lname)
	for i, v := range p1.flavour{
		fmt.Println(i, v)	
	}
	fmt.Println(p2.fname)
	fmt.Println(p2.lname)
	for i, v := range p2.flavour{
		fmt.Println(i, v)	
	}
}
