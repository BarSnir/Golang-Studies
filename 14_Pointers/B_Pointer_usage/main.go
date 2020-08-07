package main

import (
	"fmt"
)

func main() {
	a := 50
	fmt.Println(a)
	// noPointerChange(a)
	pointerChange(&a)
	fmt.Println(a)
}

func noPointerChange(x int) {
	fmt.Println(x)
	x = 43
	fmt.Println(x)
}

func pointerChange(x *int) {
	*x = 43
	fmt.Println(*x)
}
