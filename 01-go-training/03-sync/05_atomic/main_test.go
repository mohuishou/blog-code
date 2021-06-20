package main

import (
	"testing"
	"time"
)

type iConfig interface {
	Get() []int
	Set([]int)
}

func bench(b *testing.B, c iConfig) {
	go func() {
		for {
			time.Sleep(1000 * time.Nanosecond)
			c.Set([]int{100})
		}
	}()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
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

func BenchmarkConfig(b *testing.B) {
	conf := &Config{}
	conf.Set([]int{})
	bench(b, conf)
}
