package main

import (
	"fmt"
)

func main() {
	// for init; while; post
	for i := 0; i < 100; i++ {
		fmt.Println("Hello loops", i)
		for j := 0; j < 3; j++ {
			fmt.Println("The outer loop", i, "The inner loop", j)
		}
	}
}
