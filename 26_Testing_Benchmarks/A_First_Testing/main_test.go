package main

import (
	"testing"
)

func TestMySum(t *testing.T) {
	x := mySum(1, 4)
	if x != 5 {
		t.Error("Expected", 5, "Got", x)
	}
	// for 100% coverage
	main()
}
