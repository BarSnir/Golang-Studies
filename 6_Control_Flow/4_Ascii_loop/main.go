package main

import (
	"fmt"
)

func main() {
	// for init; while; post
	for i := 33; i <= 222; i++ {
		fmt.Printf("%v\n %#U\n %#x\n", i, i, i)
	}
}
