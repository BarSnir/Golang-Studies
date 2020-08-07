package main

import (
	"fmt"
	"sort"
)

func (p ByAge) Len() int           { return len(p) }
func (p ByAge) Less(i, j int) bool { return p[i].age < p[j].age }
func (p ByAge) Swap(i, j int)      { p[i].age, p[j] = p[j].age, p[i] }

// Person for main struct
type Person struct {
	first string
	age   int
}

// ByAge for sort
type ByAge []Person

func main() {
	p1 := Person{
		first: "Bar",
		age:   29,
	}

	p2 := Person{
		first: "Meshi",
		age:   26,
	}

	pepole := []Person{p1, p2}
	fmt.Println(pepole)
	sort.Sort(ByAge(pepole))
	fmt.Println(pepole)
}
