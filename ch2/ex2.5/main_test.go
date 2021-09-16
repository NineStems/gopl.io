package main

import (
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountOLD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountOld(0x1234567890ABCDEF)
	}
}



