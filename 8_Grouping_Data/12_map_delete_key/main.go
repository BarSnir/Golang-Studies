package main

import "fmt"

func main() {
	m := map[string]int{
		"James": 32,
		"Miss":  27,
	}
	m["todd"] = 33
	delete(m, "todd")
	for i, v := range m {
		fmt.Println(i, v)
	}
}
