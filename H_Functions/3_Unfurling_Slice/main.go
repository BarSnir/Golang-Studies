package main

import (
	"fmt"
)

func main() {
	xi := []int{2, 3, 4, 5}
	sum := foo(xi...)
	sum2 := foo()
	fmt.Println("The total is", sum)
	fmt.Println("The of unfurling slice is", sum2)
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
