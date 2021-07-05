package main

import "fmt"

func main() {
	//f := bar()
	//fmt.Printf("%T\n", f) //type of  : func() int
	//i := f()
	//fmt.Println(i)
	// above code or
	fmt.Println(bar()())

}

//returning a func
func bar() func() int {
	fmt.Println("I am returning a function")
	return func() int {
		return 42
	}
}
