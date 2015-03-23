package main

import (
	"github.com/funkygao/gobench/util"
	"sync"
	"testing"
)

func main() {
	b := testing.Benchmark(benchmarkDefer)
	util.ShowBenchResult("defer", b)
	b = testing.Benchmark(benchmarkDeferUnlock)
	util.ShowBenchResult("defer mutex unlock", b)
	b = testing.Benchmark(benchmarkNodeferUnlock)
	util.ShowBenchResult("no defer mutex unlock", b)
}

func benchmarkDefer(b *testing.B) {
	b.ReportAllocs()
	f := func() {
		defer func() {
		}()
	}

	for i := 0; i < b.N; i++ {
		f()
	}
}

func benchmarkDeferUnlock(b *testing.B) {
	b.ReportAllocs()
	var mu sync.Mutex
	for i := 0; i < b.N; i++ {
		deferUnlock(mu)
	}
}

func benchmarkNodeferUnlock(b *testing.B) {
	b.ReportAllocs()
	var mu sync.Mutex
	for i := 0; i < b.N; i++ {
		nodeferUnlock(mu)
	}
}

func deferUnlock(mu sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
}

func nodeferUnlock(mu sync.Mutex) {
	mu.Lock()
	mu.Unlock()
}
