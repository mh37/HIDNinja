package main

import (
	"testing"
)

func BenchmarkTranslationLayer(b *testing.B) {
	payload := "hello world this is a test payload to check performance with lowercase letters 0123456789!@#$"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, ch := range payload {
			_, keyByte := charToKeystroke(ch)
			runeKey := rune(keyByte)
			_ = translationLayer(runeKey)
		}
	}
}
