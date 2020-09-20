package mymath

import "fmt"

func sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	fmt.Println(sum)
	return sum
}
