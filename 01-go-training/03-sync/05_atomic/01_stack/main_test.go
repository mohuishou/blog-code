package main

import (
	"testing"
)

func BenchmarkNewLFStack(b *testing.B) {
	s := NewLFStack()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			s.Push(1)
		}
	})
}
