package main

import (
	"fmt"
)

func main() {
	myIncrementor := incrementor()
	fmt.Println(myIncrementor())
	fmt.Println(myIncrementor())
	fmt.Println(myIncrementor())
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
