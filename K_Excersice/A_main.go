package main

import "fmt"

func main() {
	p1 := struct {
		first string
		last  string
		flavors []string
	}{
		first: "Bar",
		last:  "Snir",
		flavors: []string{"Chocolate, Vanila"},
	}
	fmt.Println(p1)
}

// Aggregate data structure.