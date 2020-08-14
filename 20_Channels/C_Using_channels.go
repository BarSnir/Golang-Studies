package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go foo(c)
	//wating for what is does, no needed WaitGroup
	bar(c)
	fmt.Println("The program about to Exit!")
}

func foo(c chan<- int) {
	c <- 42
}

func bar(c <-chan int) {
	fmt.Println("From channel c", <-c)
}
