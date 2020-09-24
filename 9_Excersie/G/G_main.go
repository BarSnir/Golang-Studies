package main

import "fmt"

func main() {
	lettersEmojis := map[string]string{
		"A": "😈",
		"B": "😎",
		"C": "😏",
	}
	fmt.Println(lettersEmojis)
	fmt.Printf("%T\n", lettersEmojis)
}
