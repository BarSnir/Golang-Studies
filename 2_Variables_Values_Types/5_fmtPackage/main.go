package main

import (
	"fmt"
)

var y int = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y);
	fmt.Printf("%b\n", y);
	fmt.Printf("%x\n", y);
	fmt.Printf("%#x\n", y);
	y = 911;

	// can assign to a variable.
	x := fmt.Sprintf("%b\n", y);
	fmt.Printf("%#x\n%b\n", y, y);
	s := fmt.Sprintf("%#x\n%b\n", y, y);
	fmt.Println(s);
	fmt.Println(x);

	// default format
	fmt.Printf("%v", y)
}
