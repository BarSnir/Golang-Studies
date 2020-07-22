package main

import (
	"fmt"
)

func main() {
	agent := "Smith"

	switch agent {
	case "Smith":
		fmt.Println("Agent", agent, "is here!")
	case "false":
		fmt.Println("Who is this noob agent", agent, "?")
	}
}
