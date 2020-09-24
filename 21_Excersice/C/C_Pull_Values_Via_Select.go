package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
		q <- 1
	}()
	return c
}

func receive(c, q <-chan int) {
	for {
		select {
		case v, ok := <-c:
			fmt.Println("Channel C", v)
			fmt.Println("Channel C", ok)
		case v := <-q:
			fmt.Println("Channel Quit", v)
			return
		}
	}

}
