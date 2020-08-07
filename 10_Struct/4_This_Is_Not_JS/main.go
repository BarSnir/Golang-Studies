package main

import "fmt"

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "Bar",
		last:  "Snir",
		age:   29,
	}
	p2 := p1
	p2.first = "Meshi"
	p2.first = "Mor"
	p2.age = 26
	fmt.Println(p2)
	fmt.Println(p1)
}

// Aggregate data structure.
