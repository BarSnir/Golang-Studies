package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println(a)
	// This is the pointer VV
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)
	b := &a
	// Convert memory slot value
	fmt.Println(*b)

	*b = 40
	// Change the pointer value
	fmt.Println(a)
}
