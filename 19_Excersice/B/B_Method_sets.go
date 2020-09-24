package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type human interface {
	speak()
}

func (person *person) speak() {
	fmt.Println(person.first)
}

func main() {
	p1 := person{
		first: "Bar",
		last:  "Snir-Schlinger",
		age:   29,
	}
	saySomthing(&p1)
}

func saySomthing(human human) {
	human.speak()
}
