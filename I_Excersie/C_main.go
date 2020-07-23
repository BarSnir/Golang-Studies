package main

import "fmt"

func main() {
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	sliceA := slice[:5]
	fmt.Println(sliceA)
	sliceB := slice[5:]
	fmt.Println(sliceB)
	sliceC := slice[2:7]
	fmt.Println(sliceC)
	sliceD := slice[1:6]
	fmt.Println(sliceD)
}
