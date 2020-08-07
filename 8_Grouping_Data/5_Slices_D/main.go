package main

import "fmt"

func main() {
	// type{values} Composite Literal
	x := []int{4, 5, 6}
	x = append(x, 7, 8, 9) // variodic params
	fmt.Println(x)
}

// A Slice is allow you to group values of the same type
