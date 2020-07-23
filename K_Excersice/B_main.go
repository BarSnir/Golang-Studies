package main

import "fmt"

type person struct {
	first   string
	last    string
	flavors []string
}

func main() {
	p1 := person{
		first:   "Bar",
		last:    "Snir",
		flavors: []string{"Chocolate", "Vanila"},
	}
	p2 := person{
		first:   "Meshi",
		last:    "Mor",
		flavors: []string{"Cream", "cookies"},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(m)
}

// Aggregate data structure.
