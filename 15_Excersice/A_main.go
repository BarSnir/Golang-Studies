package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: "Bar",
		last:  "Snir",
	}
	noChange(p1)
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person) {
	p.last = "Schlinger"
}

func noChange(p person) {
	p.last = "Snir"
}
