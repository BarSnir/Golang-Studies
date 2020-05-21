package main

import (
	"fmt"
)

func main() {
	// for init; while; post
	x := 1
	y := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}
	for {
		if y > 8 {
			break
		}
		fmt.Println(y)
		y++
	}
	fmt.Println("Done.")
}
