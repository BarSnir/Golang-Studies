package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)
	cr := make(chan<- int, 1) // receive
	cs := make(<-chan int, 1) // send

	go func() {
		c <- 42
		cr <- (<-c)
	}()
	// fmt.Println(<-cr) wont work because the channel's type is receive only.

	fmt.Println(<-c)
	fmt.Printf("%T\t", cs)
	fmt.Printf("%T\t", cr)
}
