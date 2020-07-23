package main

import "fmt"

func main() {
	lettersEmojis := map[string]string{
		"A": "😈",
		"B": "😎",
		"C": "😏",
	}

	delete(lettersEmojis, "A")
	for i, v := range lettersEmojis {
		fmt.Println(i, v)
	}
}
