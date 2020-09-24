package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	fmt.Println("Goroutines Num:", runtime.NumGoroutine())
	fmt.Println("Hello from main!")
	wg.Wait()
}

func foo() {
	fmt.Println("Hello from foo!")
	wg.Done()
}

func bar() {
	fmt.Println("Hello from bar!")
	wg.Done()
}
