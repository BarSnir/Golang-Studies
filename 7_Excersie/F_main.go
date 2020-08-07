package main

import (
	"fmt"
)

func main() {
	agent := "Smith"
	if agent == "Bond" {
		fmt.Println("We dead, agent", agent, "is here!")
	} else {
		fmt.Println("Who is this noob agent", agent, "?")
	}
}
