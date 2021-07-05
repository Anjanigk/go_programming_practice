package main

import "fmt"

//closure: limiting the scope of a variable

func main() {
	x := incrementMe()
	fmt.Println(x()) //1
	fmt.Println(x()) //2
	fmt.Println(x()) //3
	fmt.Println(incrementMe()()) //1

}
func incrementMe() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
