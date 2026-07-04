package main

import (
	"testing"
)

func BenchmarkTranslationLayer(b *testing.B) {
	payload := "HELLO WORLD THIS IS A TEST PAYLOAD TO CHECK PERFORMANCE"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, ch := range payload {
			_, keyStr := charToKeystroke(ch)
			var runeKey rune
			if len(keyStr) > 0 {
				runeKey = rune(keyStr[0])
			}
			_ = translationLayer(runeKey)
		}
	}
}
