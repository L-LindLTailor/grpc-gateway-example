package bench

import "testing"

func BenchmarkA(b *testing.B) { // TODO
	for i := 0; i < b.N; i++ {
	}
}

func BenchmarkB(b *testing.B) { // TODO
	for i := 0; i < b.N; i++ {
	}
}
