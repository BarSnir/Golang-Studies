package main

import (
	"fmt"
)

func main() {
	// for init; while; post
	foo()
	bar("Bar")
	s1 := woo("Bar")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x, y)
}

// func
//func (r receiver) identifier(params) (return(s)) {...}

func foo() {
	fmt.Println("Hello from foo")
}

func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from foo to", s)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello`)
	b := false
	return a, b
}
