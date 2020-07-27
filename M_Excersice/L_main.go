package main

import (
	"fmt"
)

func main() {
	evalRecursion(4)
}

func evalRecursion(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Println(n)
	return evalRecursion(n - 1)
}
