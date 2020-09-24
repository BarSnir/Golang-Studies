package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var incrementor int64
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
	atomic.AddInt64(&incrementor, 1)
	// Run it with --race will sync this increment race.
	runtime.Gosched()
	fmt.Println("Incrementor counts:", atomic.LoadInt64(&incrementor))
	wg.Done()
}
