package main

import "fmt"

func main() {
	for i := 32 ; i <= 122 ; i++ {
		fmt.Printf("%v\t%#U\n", i,i)
	}
}
