package main

import (
	"fmt"
)

type bar int

var x bar = 5

func main() {
	s := int(x)
	fmt.Printf("%T", s)
}
