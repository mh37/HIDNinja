package main

import (
	"strings"
	"testing"
)

func BenchmarkToUpper(b *testing.B) {
	ch := rune('a')
	for i := 0; i < b.N; i++ {
		_ = strings.ToUpper(string(ch))
	}
}

func BenchmarkRuneMath(b *testing.B) {
	ch := rune('a')
	for i := 0; i < b.N; i++ {
		_ = string(ch - 32)
	}
}
