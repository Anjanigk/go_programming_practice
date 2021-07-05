package main

import "fmt"

//callback function: func as an argument

func main() {
	ir := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Sum of all numbers", sum(ir...))
	fmt.Println("Sum of odd numbers", oddSum(sum, ir...))

}
func sum(i ...int) int {
	s := 0
	for _, v := range i {
		s += v
	}
	return s
}

//func as an argument
func oddSum(f func(i ...int) int, n ...int) int {
	l := []int{}
	for _, v := range n {
		if v%2 != 0 {
			l = append(l, v)
		}
	}
	return f(l...)
}
