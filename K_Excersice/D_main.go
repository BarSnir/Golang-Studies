package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

func main() {
	t1 := struct {
		vehicle
		fourWheel bool
	}{
		vehicle: vehicle{
			doors: 7,
			color: "Red",
		},
		fourWheel: true,
	}

	fmt.Println(t1)

	s1 := struct {
		vehicle
		luxury bool
	}{
		vehicle: vehicle{
			doors: 5,
			color: "Black",
		},
		luxury: true,
	}

	fmt.Println(s1)
}

// Aggregate data structure.
