package main

import "fmt"
// func using variadic parameter and slice
func main() {
	sum1 := foo([]int{1, 2, 3, 4, 5}...)
	sum2 := bar([]int{1, 2, 3, 4, 5})
	fmt.Println("sum1 :", sum1)
	fmt.Println("sum2 :", sum2)
}

func foo(p ...int) int {
	sum := 0
	for _, v := range p {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
