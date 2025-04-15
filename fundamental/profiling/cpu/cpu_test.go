package cpu

import "testing"

const (
	iterations = 1000
	testString = "abc"
)

func BenchmarkConcatInefficient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatInefficient(iterations, testString)
	}
}

func BenchmarkConcatEfficient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatEfficient(iterations, testString)
	}
}
