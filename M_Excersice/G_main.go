package main

import (
	"fmt"
)

func main() {
	myFunc := func() {
		fmt.Println("Hello world!")
	}
	myFunc()
	fmt.Printf("%T\n", myFunc)
}
