package main

import "fmt"

func main() {
	total := foo(1, 2, 3)
	fmt.Println(total)

	slice := []int{1, 2, 4}
	totalBar := bar(slice)

	defer func() {
		fmt.Println(totalBar)
	}()

	fmt.Println("First this print")
}

func foo(y ...int) int {
	var sum int
	for _, v := range y {
		sum += v
	}
	return sum
}

func bar(z []int) int {
	var sum int
	for _, v := range z {
		sum += v
	}
	return sum
}
