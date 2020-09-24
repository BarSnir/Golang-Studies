package main

import (
	"fmt"
)

var x int = 1
var y string = "Hello"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	fmt.Println(s)
}
