package main

import "fmt"

//func expression
func main() {
	f := func() int {
		fmt.Println("I am func expression")
		return 10
	}
	i := f()
	fmt.Println(i)
	g := func(i int) {
		fmt.Println("I am func expression with input", i)
	}
	g(42)
}
