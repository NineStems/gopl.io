package main

import "fmt"

func main() {
	x := uint64(10)
	fmt.Println(PopCount(x))
	fmt.Println(PopCountOld(x))

}

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var res int
	mask := uint64(0b1)
	for i := uint64(0); i < 64; i++ {
		if mask & (x >> i) > 0 {
			res++
		}
	}
	return res

}

func PopCountOld(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
