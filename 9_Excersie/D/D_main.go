package main

import "fmt"

func main() {
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	sliceY := []int{56, 57, 58, 59, 60}
	slice = append(slice, sliceY...)
	fmt.Println(slice)
}
