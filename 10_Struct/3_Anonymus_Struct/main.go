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
	fmt.Println(p1)
}

// Aggregate data structure.
