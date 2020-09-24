package main

import "fmt"

func main() {
	c := make(chan int, 1)
	d := make(chan int, 2)
	// This is a deadlock, it will fail.
	//c <- 42
	//fmt.Println(<-c)

	go func() {
		c <- 42
		d <- 44
		d <- 45
		// c <- 43 is a unsuccessfully buffer. Also an deadlock.
	}()
	fmt.Println(<-c)
	fmt.Println(<-d)
	fmt.Println(<-d)
}
