package main

import (
	"fmt"
)

func main() {
	myFunc := func() {
		fmt.Println("Hello Ravit")
	}
	evalFunc(myFunc)
}

func evalFunc(myFunc func()) {
	myFunc()
}
