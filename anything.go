package main

import "fmt"

func main() {
	fmt.Println("First programme")
	foo()
	fmt.Println("something more")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("I am in foo")
}

func bar() {
	fmt.Println("Exiting!!")
}
