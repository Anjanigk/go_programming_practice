package main

import "fmt"

//recursion

func main() {
	fmt.Println(factorial(4)) //24
	fmt.Println(loopMult(4))  // 24

}
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

//via loop
func loopMult(n int) int {
	res := 1
	for ; n > 0; n-- {
		res *= n
	}
	return res
}
