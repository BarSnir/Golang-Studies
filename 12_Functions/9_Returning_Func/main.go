package main

import (
	"fmt"
)

func main() {
	s := foo()
	fmt.Println(s)
	i := bar()
	fmt.Printf("%T\n", bar)
	fmt.Printf("%T\n", i)
	fmt.Println(i())
}

func foo() string {
	return "Hello world"
}

func bar() func() int {
	// with closure ;)
	x := 4
	return func() int {
		return x
	}
}
