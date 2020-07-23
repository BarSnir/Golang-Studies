package main

import "fmt"

func main() {
	// type{values} Composite Literal
	x := []int{4, 5, 6}
	fmt.Println(x[0:])
	fmt.Println(x[1:2])
	fmt.Println(x[0:2])
	fmt.Println(x[1:3])
}

// A Slice is allow you to group values of the same type
