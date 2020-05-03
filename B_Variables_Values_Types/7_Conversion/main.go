package main

import (
	"fmt"
)

var a int = 42

type hotdog int

var b hotdog = 52

func main() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// (type of a != type of b)!

	a = int(b)
	fmt.Println(a)
}
