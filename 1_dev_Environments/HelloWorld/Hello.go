package hello

import (
	"fmt"

	"rsc.io/quote"
)

func hello() {
	fmt.Println("Hello world.")
}

func helloB() string {
	return quote.Hello()
}
