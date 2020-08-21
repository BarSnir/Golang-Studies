package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	routines := 10
	for i := 0; i < routines; i++ {
		wg.Add(1)
		go func(i int) {
			for j := i * routines; j < (i*routines)+routines; j++ {
				ch <- j
			}
			wg.Done()
		}(i)
	}

	for j := 0; j < routines*routines; j++ {
		fmt.Println(<-ch)
	}

	// for v := range ch {
	// 	fmt.Println(v)
	// }
	wg.Wait()
}
