package main

import "fmt"

//slice of slice
func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x[:5])  //[42 43 44 45 46]
	fmt.Println(x[5:])  //[47 48 49 50 51]
	fmt.Println(x[2:7]) //[44 45 46 47 48]
	fmt.Println(x[1:6]) //[43 44 45 46 47]

}
