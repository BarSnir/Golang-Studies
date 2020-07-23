package main

import (
	"fmt"
)

func main() {
	// closing the file with defer
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
