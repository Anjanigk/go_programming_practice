package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	p.first = "John"
	//(*p).first = "Sam"
	p.last = "Smith"
	p.age = 50
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   42,
	}

	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}
