package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"age"`
}

func main() {
	str := `[
		{"first":"Bar","last":"Snir",
		 "age":29},{"first":"Meshi","last":"Mor", "age":26}
	]`
	// byte slice
	bs := []byte(str)

	var pepole []person
	err := json.Unmarshal(bs, &pepole)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pepole)

	for _, value := range pepole {
		fmt.Println(value.First)
	}
}
