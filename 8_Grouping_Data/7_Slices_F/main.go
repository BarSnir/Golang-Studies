package main

import "fmt"

func main() {
	// type{values} Composite Literal
	x := []int{4, 5, 6}
	x = append(x, 7, 8, 9) // variodic params
	fmt.Println(x)

	y := []int{200, 300, 400}
	x = append(x, y...)
	fmt.Println("After appending variodic params", x)

	// to the first & second item, append from the fourth item, all.
	x = append(x[:2], x[4:]...)

	fmt.Println("After delete with append", x)
}

// A Slice is allow you to group values of the same type
