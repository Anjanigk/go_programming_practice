package main

import "fmt"

func main() {
	favSport := "karate"
	switch favSport {
	case "cycling":
		fmt.Println("I wont execute.")
	case "karate":
		fmt.Println("I will execute")
	default:
		fmt.Println("I am default")
	}
}
