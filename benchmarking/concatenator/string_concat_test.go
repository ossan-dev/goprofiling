package concatenator_test

import (
	"testing"

	"benchmarking/concatenator"
)

var result string // avoid compiler optimizations

func BenchmarkStrings(b *testing.B) {
	var s string
	b.ReportAllocs()
	b.ResetTimer() // you can reset timers/counters for heavy setup
	for i := 0; i < b.N; i++ {
		s = concatenator.Strings("hello", "world")
	}
	result = s // use a pkg level var
}
