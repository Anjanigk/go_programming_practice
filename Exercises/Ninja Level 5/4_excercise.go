package main

import "fmt"

//Anonymous Struct
func main() {
	p := struct {
		fname   string
		lname   string
		flavour []string
	}{
		fname:   "John",
		lname:   "Smith",
		flavour: []string{"Vanila", "Choclate", "Butterscoch"},
	}

	fmt.Println(p)

}
