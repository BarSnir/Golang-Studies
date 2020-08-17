package main

import (
	"fmt"
	"sync"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)
	go send(eve, odd)
	go receive(eve, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}
	fmt.Println("The program about to exit!")
}

func send(eve, odd chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			eve <- i
		} else {
			odd <- i
		}
	}
	close(eve)
	close(odd)
}

func receive(e, o <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range e {
			fanin <- v

		}
		close(fanin)
		wg.Done()
	}()

	go func() {
		for v := range o {
			fanin <- v
		}
		close(fanin)
		wg.Done()
	}()
}
