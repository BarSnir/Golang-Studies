package main

import "testing"

func TestMyMain(t *testing.T) {
	main()
}

func BenchmarkMyMain(b *testing.B) {
	main()
}
