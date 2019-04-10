package main

import (
	"testing"
)

func BenchmarkWordCounter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wordCounter()
	}
}

func BenchmarkWordCounterConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wordCounterConcurrent()
	}
}
