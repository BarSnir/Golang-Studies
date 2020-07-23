package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("Anonymus function is here.")
	}()
	func(x int) {
		fmt.Println(x)
	}(42)
}
