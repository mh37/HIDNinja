package main

import (
	"os"
	"testing"
)

func BenchmarkSendKeyOld(b *testing.B) {
	// create dummy /dev/hidg0 locally or just a temp file
	tmpfile, _ := os.CreateTemp("", "hidg0")
	defer os.Remove(tmpfile.Name())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f, _ := os.OpenFile(tmpfile.Name(), os.O_APPEND|os.O_WRONLY, 0666)
		f.Write([]byte{0x00})
		f.Close()
	}
}

func BenchmarkSendKeyNew(b *testing.B) {
	tmpfile, _ := os.CreateTemp("", "hidg0")
	defer os.Remove(tmpfile.Name())

	f, _ := os.OpenFile(tmpfile.Name(), os.O_APPEND|os.O_WRONLY, 0666)
	defer f.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Write([]byte{0x00})
	}
}
