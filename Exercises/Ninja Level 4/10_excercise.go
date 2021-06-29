package main

import "fmt"

func main() {
	m := map[string][]string{
		"james":      []string{`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny": []string{`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`},
		"no_dr":      []string{`no_dr`, `Being evil`, `Ice cream`, `Sunsets`},
	}
	if _, ok := m["moneypenny"]; ok {
		//fmt.Println("value:",i)
		delete(m, "moneypenny")
	}
	for key, values := range m {
		fmt.Println("this is a record for:", key)
		for i, v := range values {
			fmt.Println("\t", i, v)
		}
	}

}
