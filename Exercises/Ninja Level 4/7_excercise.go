package main

import "fmt"
func main() {
	x := []string{"James", "Bond", "Shaken, not stirred"}
	fmt.Println(x)
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(y)
	ms := [][]string {x,y}
	fmt.Println(ms)
}