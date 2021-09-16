package main

import (
	"gopl.io/ch2/popcount"
	"testing"
)

func BitCount(x uint64) int {
	// Hacker's Delight, Figure 5-2.
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}

func PopCountByClearing(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		n++
	}
	return n
}

func PopCountByShifting(x uint64) int {
	n := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			n++
		}
	}
	return n
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkBitCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(0x1234567890ABCDEF)
	}
}


/*
old
goos: linux
goarch: amd64
pkg: gopl.io/ch2/popcount
cpu: Intel(R) Core(TM) i5-7600K CPU @ 3.80GHz
BenchmarkPopCount
BenchmarkPopCount-4             	1000000000	         0.2560 ns/op
BenchmarkBitCount
BenchmarkBitCount-4             	1000000000	         0.2581 ns/op
BenchmarkPopCountByClearing
BenchmarkPopCountByClearing-4   	100000000	        10.66 ns/op
BenchmarkPopCountByShifting
BenchmarkPopCountByShifting-4   	69472080	        16.80 ns/op
PASS
new
goos: linux
goarch: amd64
pkg: gopl.io/ch2/ex2.3
cpu: Intel(R) Core(TM) i5-7600K CPU @ 3.80GHz
BenchmarkPopCount
BenchmarkPopCount-4             	1000000000	         0.2507 ns/op
BenchmarkBitCount
BenchmarkBitCount-4             	1000000000	         0.2499 ns/op
BenchmarkPopCountByClearing
BenchmarkPopCountByClearing-4   	100000000	        10.64 ns/op
BenchmarkPopCountByShifting
BenchmarkPopCountByShifting-4   	69206107	        16.88 ns/op
PASS
 */


