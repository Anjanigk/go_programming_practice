package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)  //value
	fmt.Println(&a) //address
	b := &a
	fmt.Println(b)   //address
	fmt.Println(*b)  //value
	fmt.Println(*&a) //value

	*b = 43        //value changed of a
	fmt.Println(a) //43
}
