package main

import "fmt"

func main() {
	// type{values} Composite Literal
	jb := []string{"james", "bond", "chocolate", "martini"}
	mp := []string{"miss", "Money", "Strawberry", "vodka"}

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}

// A Slice is allow you to group values of the same type
