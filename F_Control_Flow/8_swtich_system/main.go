package main

import (
	"fmt"
)

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print")
	case (2 == 2):
		fmt.Println("this should print")
		fallthrough
	case (4 == 4):
		fmt.Println("and this should print")
	}

	n := "d"
	switch n {
	case "Bond":
		fmt.Println("Bond, james Bond.")
	case "a", "b", "c":
		fmt.Println("A, B, C")
	default:
		fmt.Println("Default")
	}

}
