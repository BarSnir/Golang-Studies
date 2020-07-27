package main

import (
	"fmt"
)

func main() {
	strFunc := generateFunc()
	str := strFunc()
	fmt.Println(str)
}

func generateFunc() func() string {
	return func() string{
		return "Hello Bar"
	}
}