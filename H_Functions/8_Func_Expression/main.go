package main

import (
	"fmt"
)

func main() {
	f := func(){
		fmt.Println("My first experssion.")
	}
	// first class citizen
	f()
}