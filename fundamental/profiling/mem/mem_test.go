package mem

import "testing"

const dataCount = 5000

func BenchmarkProcessDataInefficient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = processDataInefficient(dataCount)
	}
}

func BenchmarkProcessDataEfficient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = processDataEfficient(dataCount)
	}
}
