package main

import "fmt"
// multiple return from a function
func main() {
	i := foo()
	j, s := bar()
	fmt.Println(i)
	fmt.Println(j, s)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 89, "The End!"
}
