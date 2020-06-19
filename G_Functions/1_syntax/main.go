package main

import (
	"fmt"
)

func main() {
	// for init; while; post
	foo()
	bar("Bar")
}

// func
//func (r receiver) identifier(params) (return(s)) {...}

func foo() {
	fmt.Println("Hello from foo")
}

func bar(s string) {
	fmt.Println("Hello,", s)
}
