package main

import (
	"testing"
)

type iConfig interface {
	Get() []int
	Set([]int)
}

func bench(b *testing.B, c iConfig) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			c.Set([]int{100})
			c.Get()
			c.Get()
			c.Get()
			c.Set([]int{100})
			c.Get()
			c.Get()
		}
	})
}

func BenchmarkMutexConfig(b *testing.B) {
	conf := &MutexConfig{data: []int{1, 2, 3}}
	bench(b, conf)
}

func BenchmarkRWMutexConfig(b *testing.B) {
	conf := &RWMutexConfig{data: []int{1, 2, 3}}
	bench(b, conf)
}
