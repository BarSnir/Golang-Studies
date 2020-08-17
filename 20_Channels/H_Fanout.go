package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go thortelingRoutines(c1, c2)

	for myValue := range c2 {
		fmt.Println(myValue)
	}

	fmt.Println("The program about to exit!")
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func(v2 int) {
			c2 <- timeConsumeWork(v2)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)
}

func timeConsumeWork(n int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}

func thortelingRoutines(c1, c2 chan int) {
	var wg sync.WaitGroup
	const goroutines = 10

	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		for v := range c1 {

			go func(v2 int) {
				c2 <- timeConsumeWork(v2)
				wg.Done()
			}(v)
		}
	}
	wg.Wait()
	close(c2)
}
