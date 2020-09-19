package main

import (
	"fmt"
)

type customErr struct{
	err string
}

func (ce customErr) Error() string {
	return "this is my error message"
}

func main()  {
	c1 := customErr{}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}