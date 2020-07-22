package main

import (
	"fmt"
)

func main() {
	agent := "Smith"

	switch {
	case true:
		fmt.Println("We dead, agent", agent, "is here!")
	case false:
		fmt.Println("Who is this noob agent", agent, "?")
	}
}
