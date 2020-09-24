package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// run only with flag --race

func main() {

	var counter int64
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			// Mutex for Prevent race condition.
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter \t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter \t", counter)
}
