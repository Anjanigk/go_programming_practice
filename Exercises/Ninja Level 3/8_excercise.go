package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("I wont execute.")
	case 2 == 2:
		fmt.Println("I will execute")
	default:
		fmt.Println("I am default")
	}
}
