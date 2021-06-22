package main

import "fmt"

func main() {
	n, e := fmt.Println("First programme", 12, true)
	//foo()
	//fmt.Println("something more")
	//for i := 0; i < 100; i++ {
	//	if i%2 == 0 {
	//		fmt.Println(i)
	//	}
	//}
	//bar()
	fmt.Println(n)
	fmt.Println(e)
}

func foo() {
	fmt.Println("I am in foo")
}

func bar() {
	fmt.Println("Exiting!!")
}
