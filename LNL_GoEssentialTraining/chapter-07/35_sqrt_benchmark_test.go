package main

import (
	"testing"
)

func almostEqualAgainAgain(v1, v2 float64) bool {
	return Abs(v1-v2) <= 0.001
}

func TestSimpleAgain(t *testing.T) {
	val, err := Sqrt(2)

	if err != nil {
		t.Fatalf("Error %s", err)
	}

	if !almostEqualAgainAgain(val, 1.414214) {
		t.Fatalf("bad value %f", val)
	}
}

func BenchmarkSqrt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Sqrt(float64(i))
		if err != nil {
			b.Fatal(err)
		}
	}
}
