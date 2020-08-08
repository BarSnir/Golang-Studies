package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	// runs before main.
}

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPU\t", runtime.NumCPU())
	wg.Add(1)
	go foo()
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo: ", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar: ", i)
	}
}