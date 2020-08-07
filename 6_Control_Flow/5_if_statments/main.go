package main

import (
	"fmt"
)

func main() {

	if true {
		fmt.Println("True-A")
	}
	if false {
		fmt.Println("True-B")
	}
	if !true {
		fmt.Println("True-C")
	}
	if !false {
		fmt.Println("True-D")
	}
	if !(2 == 2) {
		fmt.Println("True-E")
	}
	if !(2 != 2) {
		fmt.Println("True-F")
	}
	if x := 2; x == 2 {
		fmt.Println("True-G")
	}
}
