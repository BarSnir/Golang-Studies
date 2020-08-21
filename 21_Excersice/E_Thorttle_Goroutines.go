package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {

			for j := i * 10; j < (i*10)+10; j++ {
				c <- j
			}
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
