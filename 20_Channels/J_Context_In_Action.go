package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("error check 1:", ctx.Err())
	fmt.Println("num gortins 1:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			fmt.Println("num gortins:", runtime.NumGoroutine())
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 20)
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(time.Second * 3)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("num goroutine 2:", runtime.NumGoroutine())

	fmt.Println("about to cancel context")
	cancel()
	fmt.Println("cancelled context")

	time.Sleep(time.Second * 2)
	fmt.Println("error check 3:", ctx.Err())
	fmt.Println("num goroutine 3:", runtime.NumGoroutine())
}
