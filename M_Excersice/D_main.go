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
		last:  "Foo",
		age:   29,
	}
	p1.speak()
}

func (p person) speak() {
	fmt.Println("I'm", p.first, p.last)
}