package main

import (
	"fmt"
	"runtime"
)

var x int
var y float64
var limit int8

func main() {
	x = 5
	y = 3.545324
	// limit = 600 error
	limit = 127
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\t", x)
	fmt.Printf("%T\n", y)
	fmt.Println(limit)

	// some printing from runtime
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
