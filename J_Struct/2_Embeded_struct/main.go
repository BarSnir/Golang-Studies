package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	p1 := person{
		first: "Bar",
		last:  "Snir",
		age:   29,
	}

	agent := secretAgent{
		person: p1,
		ltk: true,
	}
	fmt.Println(agent)
	fmt.Println(agent.first)
	fmt.Println(agent.age)
	fmt.Println(agent.ltk)
}

// Aggregate data structure.
