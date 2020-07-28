package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{5, 2, 5, 6, 40, 2, 7}
	strings := []string{"Sahar", "Bar", "Ad"}
	sort.Ints(arr)
	fmt.Println(arr)

	sort.Strings(strings)
	fmt.Println(strings)
}
