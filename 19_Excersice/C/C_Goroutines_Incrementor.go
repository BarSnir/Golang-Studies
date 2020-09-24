package main

import (
	"fmt"
	"runtime"
	"sync"
)

var incrementor int
var wg sync.WaitGroup

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
	incrementor++
	// Run with out me to get race condition.
	runtime.Gosched()
	fmt.Println("Incrementor counts:", incrementor)
	wg.Done()
}
