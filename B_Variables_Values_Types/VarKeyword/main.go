package main

import "fmt"

// outer scoped variable.
var z int
var a = `Also... 
A String`

func main() {
	var x = 12
	var str string = "Hello String"
	fmt.Println(z, str, x, a)
	// Print format;
	fmt.Printf("%T\n", z)
}
