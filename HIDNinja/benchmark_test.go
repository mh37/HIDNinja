package main

import (
	"testing"
)

func BenchmarkTranslationLayer(b *testing.B) {
	payload := "HELLO WORLD THIS IS A TEST PAYLOAD TO CHECK PERFORMANCE"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, ch := range payload {
			_ = translationLayer(ch)
		}
	}
}
