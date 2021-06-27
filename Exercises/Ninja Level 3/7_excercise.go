package main

import "fmt"

func main() {
	i := "anjani"
	if i == "anjani" {
		fmt.Println("first clause", i)
	} else if i == "dev" {
		fmt.Println("second clause", i)
	} else {
		fmt.Println("Try again!")
	}
}
