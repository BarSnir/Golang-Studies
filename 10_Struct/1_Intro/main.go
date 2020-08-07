package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Bar",
		last:  "Snir",
		age:   29,
	}
	fmt.Println(p1)
	fmt.Println(p1.first)
	fmt.Println(p1.age)

}

// Aggregate data structure.
