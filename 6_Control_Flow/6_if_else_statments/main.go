package main

import (
	"fmt"
)

func main() {
	x := 40
	if x == 40 {
		fmt.Println("Our value was 40")
	} else if x == 41 {
		fmt.Println("Our value was 41")
	} else {
		fmt.Println("Our value wasn't 40 or 41")
	}
}
