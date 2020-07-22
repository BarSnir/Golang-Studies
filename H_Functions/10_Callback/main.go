package main

import (
	"fmt"
)

func main() {
	il := []int{1, 2, 3, 4, 5, 6, 7, 8}
	sum1 := sum(2, 3, 4, 5)
	fmt.Println(sum1)

	s2 := even(sum, il...)
	fmt.Println(s2)
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
