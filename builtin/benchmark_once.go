package main

import (
	"github.com/funkygao/gobench/util"
	"sync"
	"testing"
)

var nop = func() {}

func main() {
	b := testing.Benchmark(benchmarkNop)
	util.ShowBenchResult("nop", b)
	b = testing.Benchmark(benchmarkOnceNop)
	util.ShowBenchResult("once.do(nop)", b)
}

func benchmarkNop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nop()
	}
}

func benchmarkOnceNop(b *testing.B) {
	var once sync.Once
	for i := 0; i < b.N; i++ {
		once.Do(nop)
	}
}
