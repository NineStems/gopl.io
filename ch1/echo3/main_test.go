package main

import (
	"testing"
)
//book task 1.2
func BenchmarkEchoEx1(b *testing.B) {
	ta := make([]string, 100000)
	for i:=0; i < b.N;i++{
		EchoEx1(ta)
	}
}

func BenchmarkEchoEx2(b *testing.B) {
	ta := make([]string, 100000)
	for i:=0; i < b.N;i++{
		EchoEx2(ta)
	}
}
