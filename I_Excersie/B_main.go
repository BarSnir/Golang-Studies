package main

import "fmt"

func main() {
	slice := []int{5, 40, 253, 6, 70}
	for _, v := range slice {
		fmt.Println(v)
	}
	fmt.Printf("%T\n", slice)
}
