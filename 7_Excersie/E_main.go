package main

import (
	"fmt"
)

func main() {
	x := 10
	for x < 100 {
		if x%4 == 0 {
			fmt.Println(x)
		}
		x++
	}
}
