package main

import (
	"fmt"
)

func main() {
	a := factorial(4)
	fmt.Println(a)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}