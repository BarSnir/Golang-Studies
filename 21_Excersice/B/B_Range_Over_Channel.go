package main

import (
	"fmt"
)

func main() {
	cs := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			cs <- i
		}
		close(cs)
	}()

	for v := range cs {
		fmt.Println(v)
	}

	fmt.Println("The program about to exit!")
}
