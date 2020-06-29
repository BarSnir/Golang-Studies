package main

import (
	"fmt"
)

func main() {
	sum := foo(2, 3, 4, 5)
	fmt.Println("The total is", sum)
}

// variadic params
func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("Index position", i, " and the items is", v)
	}
	return sum
}
