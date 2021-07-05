package main

import "fmt"

//Anonymous func
func main() {
	func() {
		fmt.Println("I am anonymous func")
	}()
	func(i int) {
		fmt.Println("I am anonymous func with input", i)
	}(42)
}
