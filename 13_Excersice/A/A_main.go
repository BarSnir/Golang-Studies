package main

import "fmt"

func main() {
	num := foo()
	num2, str := bar()
	fmt.Println(num, num2, str)
}

func foo() int {
	return 5
}

func bar() (int, string) {
	return 5, "hello"
}
