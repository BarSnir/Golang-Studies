package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name     string `json:"name"`
	Position string `json:"Position"`
}

func main() {
	object := `{"name":"Bar", "Position":"full develops"}`
	fmt.Println(object)
	var man person
	err := json.Unmarshal([]byte(object), &man)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(man.Name," ",man.Position)
}
