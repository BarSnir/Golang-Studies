package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck{
		vehicle: vehicle{
			doors: 7,
			color: "Red",
		},
		fourWheel: true,
	}

	fmt.Println(t1)

	s1 := sedan{
		vehicle: vehicle{
			doors: 5,
			color: "Black",
		},
		luxury: true,
	}

	fmt.Println(s1)
}

// Aggregate data structure.
