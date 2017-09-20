package requestid

import (
	"fmt"
	"testing"
)

func TestAtomicUUID(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(atomicUUID())
	}
}

func BenchmarkUUIDGoogle(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		googleUUID()
	}
}

func BenchmarkUUIDSatori(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		satoriUUID()
	}
}

func BenchmarkUUIDAtomic(b *testing.B) {
	for n := 0; n < b.N; n++ {
		atomicUUID()
	}
}

func BenchmarkCounterTime(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		timeCounter()
	}
}

func BenchmarkCounterAtomic(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		atomicCounter()
	}
}
