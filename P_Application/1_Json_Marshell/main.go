package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Bar",
		Last:  "Snir",
		Age:   29,
	}
	p2 := person{
		First: "Meshi",
		Last:  "Mor",
		Age:   26,
	}

	pepole := []person{p1, p2}
	fmt.Println(pepole)

	bs, err := json.Marshal(pepole)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))

}
