package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
}

func info(msg string, s shape) {
	fmt.Println(msg, s.area())
}

func main() {
	c := circle{
		radius: 12.345,
	}

	s := square{
		length: 12.65,
	}

	info("Circle area:", c)
	info("Square area:", s)
}
