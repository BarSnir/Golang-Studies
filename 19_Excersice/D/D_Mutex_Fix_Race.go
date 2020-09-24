package main

import (
	"fmt"
	"sync"
)

var incrementor int
var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	cores := 5
	wg.Add(cores)

	func() {
		for i := 0; i < cores; i++ {
			go increment()
		}
	}()
	wg.Wait()
}

func increment() {
	mu.Lock()
	incrementor++
	// Fixing the race with mutex
	fmt.Println("Incrementor counts:", incrementor)
	// run it with --race and find out that the race has gone. 
	// it will delete this block from race at all.
	mu.Unlock()
	wg.Done()
}
