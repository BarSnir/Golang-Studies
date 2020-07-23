package main

import "fmt"

func main() {
	m := map[string]int{
		"James": 32,
		"Miss":  27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Bond"]) //default to 0

	value, exists := m["Bond"]
	fmt.Println(value)
	fmt.Println(exists)

	if _, exists := m["Bond"]; exists {
		fmt.Println("The value is exists!")
	}

}
