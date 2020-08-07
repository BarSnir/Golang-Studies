package main

import (
	"fmt"
)

func main() {
	myClosure := getClosure()
	myClosure()
	myClosure()
}

func getClosure() func() {
	character := "Pirate"
	return func() {
		fmt.Println(character)
	}
}