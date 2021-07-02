package main

import "fmt"

// attaching a func with a type
//method example
type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "John",
		last:  "Smith",
		age:   25,
	}
	p2 := person{
		"James",
		"Bond",
		42,
	}
	p1.speak()
	p2.speak()
}

func (p person) speak() {
	fmt.Println("Called Speak by:", p.first, p.last, "with age:", p.age)
}
