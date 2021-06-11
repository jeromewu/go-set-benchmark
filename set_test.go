package main

import (
	"testing"
)

const it = uint(1 << 24)

func BenchmarkSetWithBoolValueWrite(b *testing.B) {
	set := make(map[uint]bool)

	for i := uint(0); i < it; i++ {
		set[i] = true
	}
}

func BenchmarkSetWithStructValueWrite(b *testing.B) {
	set := make(map[uint]struct{})

	for i := uint(0); i < it; i++ {
		set[i] = struct{}{}
	}
}

func BenchmarkSetWithInterfaceValueWrite(b *testing.B) {
	set := make(map[uint]interface{})

	for i := uint(0); i < it; i++ {
		set[i] = struct{}{}
	}
}
