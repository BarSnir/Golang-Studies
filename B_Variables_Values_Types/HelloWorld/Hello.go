package main

import "fmt"

func main() {
	fmt.Println("Hello Golang world.")
	foo()
	fmt.Println("More printing!")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func foo() {
	fmt.Println("I'm foo.")
}

func bar() {
	fmt.Println("Ant thats it. 🍻")
}

/**
Control flow:
	1) sequence.
	2) loop, iterative.
	3) conditional.
	4) when the program exit main, your proccess is over.
**/
