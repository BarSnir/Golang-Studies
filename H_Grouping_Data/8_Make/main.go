package main

import "fmt"

func main() {
	// type{values} Composite Literal
	x := make([]int, 10, 100)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[9] = 999
	fmt.Println(x)
}

// A Slice is allow you to group values of the same type
