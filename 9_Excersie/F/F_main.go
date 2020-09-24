package main

import "fmt"

func main() {
	letters := []string{"A", "B", "C"}
	emojis := []string{"😈", "😎", "😏"}
	combine := [][]string{letters, emojis}
	fmt.Println(combine)
}
