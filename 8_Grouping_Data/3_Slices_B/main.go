package main

import "fmt"

func main() {
	// type{values} Composite Literal
	x := []int{4, 5, 6}
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println(x[2])

	for i, v := range x {
		fmt.Println(i, v)
	}
}

// A Slice is allow you to group values of the same type
