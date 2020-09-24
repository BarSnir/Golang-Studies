package main

import (
	"fmt"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	go send(eve, odd, quit)
	receive(eve, odd, quit)
	fmt.Println("The program about to exit!")
}

func send(eve, odd, quit chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			eve <- i
		} else {
			odd <- i
		}
	}
	close(eve)
	close(odd)
	quit <- 0
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Channel Even", v)

		case v := <-o:
			fmt.Println("Channel Odd", v)

		case v := <-q:
			fmt.Println("Channel Quit", v)

			return
		}

	}
}
