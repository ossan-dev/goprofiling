package concatenator_test

import (
	"testing"

	"benchmarking/concatenator"
)

var Result string // "try" avoiding compiler optimizations

func BenchmarkStrings(b *testing.B) {
	var s string
	b.ReportAllocs()
	b.ResetTimer() // you can reset timers/counters for heavy setup
	for i := 0; i < b.N; i++ {
		s = concatenator.Strings("hello", "world")
	}
	Result = s // use a pkg level var
}

func BenchmarkStrings_1_24(b *testing.B) {
	b.ReportAllocs()
	// "b.Loop" tell when to stop
	for b.Loop() {
		// this won't be subject to dead code elimination
		concatenator.Strings("hello", "world")
	}
}
