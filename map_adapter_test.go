package jpostcode

import (
	"testing"
)

func Benchmark_newMapAdapter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := newMapAdapter()
		if err != nil {
			b.Error(err)
		}
	}
}
