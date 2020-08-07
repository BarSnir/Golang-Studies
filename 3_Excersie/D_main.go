package main

import (
	"fmt"
)

type bar int

var x bar = 5

func main() {
	fmt.Println(x)
	x = 42
	fmt.Println(x)
}
