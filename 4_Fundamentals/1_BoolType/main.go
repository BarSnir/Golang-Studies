package main

import (
	"fmt"
)

var isBool bool

func main() {
	fmt.Println(isBool)
	isBool = true
	fmt.Println(isBool)
	x := 5
	y := 42

	fmt.Println(x == y)
	fmt.Println(x <= y)
	fmt.Println(x >= y)
	fmt.Println(x != y)
}
