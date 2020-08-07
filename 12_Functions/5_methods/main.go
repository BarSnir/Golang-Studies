package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secertAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secertAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	sa2 := secertAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: false,
	}
	fmt.Println(sa1.person)
	fmt.Println(sa1.ltk)
	fmt.Println(sa1)

	sa1.speak()
	sa2.speak()
}

func (s secertAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}
