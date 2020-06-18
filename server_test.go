package main

import "testing"

func BenchmarkGraphql(b *testing.B) {
	// run the function b.N times
	for n := 0; n < b.N; n++ {
		main()
	}
}
