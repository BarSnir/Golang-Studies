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

type human interface {
	speak()
}

type robot struct {
	serial int
}

type robotAction interface {
	calculate()
}

func (s secertAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last)
}

func (r robot) calculate() {
	fmt.Println("Im a robot num", r)
}

func calculateIT(r robotAction) {
	fmt.Println("My serial is", r)
}

func bar(h human) {
	fmt.Println("I called human", h)
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

	p1 := person{
		first: "Dr.",
		last:  "yes",
	}

	robotA := robot{
		serial: 1,
	}

	bar(sa1)
	bar(sa2)
	bar(p1)
	calculateIT(robotA)
	robotA.calculate()
	sa1.speak()
}
