package main

import "fmt"

func main() {
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 56, 57, 58, 59, 60}
	slice = append(slice[:4], slice[10:]... )
	fmt.Println(slice)
}
