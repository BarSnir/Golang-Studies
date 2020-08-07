package main

import (
	"fmt"
)

func main() {
	s := "H"

	bs := []byte(s)
	fmt.Println(bs)

	n := bs[0]

	fmt.Println(n)
	fmt.Printf("%T\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%#X\n", n)
}
