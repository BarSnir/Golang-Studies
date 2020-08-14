package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	//wating for what is does, no needed WaitGroup
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	func() {
		//waiting for values to come. until is closed.
		for v := range c {
			fmt.Println("From channel c", v)
		}
	}()
	fmt.Println("The program about to Exit!")
}
