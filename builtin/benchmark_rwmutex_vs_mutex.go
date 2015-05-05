package main

import (
	"github.com/funkygao/gobench/util"
	"sync"
	"testing"
)

func main() {
	b := testing.Benchmark(benchmarkRWMutex)
	util.ShowBenchResult("sync.RWMutex", b)
	b = testing.Benchmark(benchmarkMutex)
	util.ShowBenchResult("sync.Mutex", b)
}

func benchmarkRWMutex(b *testing.B) {
	var m sync.RWMutex
	for i := 0; i < b.N; i++ {
		m.RLock()
		m.RUnlock()
	}
}

func benchmarkMutex(b *testing.B) {
	var m sync.Mutex
	for i := 0; i < b.N; i++ {
		m.Lock()
		m.Unlock()
	}
}
