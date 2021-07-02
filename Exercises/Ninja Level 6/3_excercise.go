package main

import "fmt"
// Use of "Defer"
func main() {
	defer foo()
	j, s := bar()
	fmt.Println(j, s)
	fmt.Println("Exiting from main")
}

func foo() {
	fmt.Println("I am in Foo!")
}

func bar() (int, string) {
	return 89, "The End!"
}
