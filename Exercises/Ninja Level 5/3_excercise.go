package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}
//Embedded structs
func main() {
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}
	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}

	fmt.Println(t)
	fmt.Println(s)
	fmt.Println("doors:", t.doors, "color:", t.color)
	fmt.Println("doors:", s.doors, "color:", s.color)

}
