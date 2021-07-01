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
		flavour: []string{"Strawberry", "Choclate"},
	}

	m := map[string]person{
		p1.lname: p1,
		p2.lname: p2,
	}
	for _, v := range m {
		fmt.Println(v.fname)
		fmt.Println(v.lname)
		for i, v := range v.flavour {
			fmt.Println(i, v)
		}
	}

}
