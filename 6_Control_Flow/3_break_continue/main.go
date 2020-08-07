package main

import (
	"fmt"
)

func main() {
	// for init; while; post
	x := 43 % 30
	y := 80 / 35
	z := 0
	fmt.Println(x, y)

	for {
		z++
		if z > 100 {
			break
		}
		if z%2 != 0 {
			continue
		}
		fmt.Println(z)

	}
}
