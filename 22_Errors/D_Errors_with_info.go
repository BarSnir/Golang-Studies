package main

import (
	"fmt"
	"log"
)

type norgateMathError struct {
	lat  string
	long string
	err  error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("norgate math occurred: %v %v", n.lat, n.long)
}

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatal(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		name := fmt.Errorf("norgate math: square root of negative number")
		return 0, norgateMathError{"59.3242 N", "99,645 W", name}
	}
	return f, nil
}
