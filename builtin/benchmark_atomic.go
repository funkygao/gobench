package main

import (
	"github.com/funkygao/gobench/util"
	"sync/atomic"
	"testing"
)

func main() {
	b := testing.Benchmark(benchmarkAddUint64)
	util.ShowBenchResult("atomic.AddUint64", b)
	b = testing.Benchmark(benchmarkAddInt32)
	util.ShowBenchResult("atomic.AddInt32", b)
}

func benchmarkAddUint64(b *testing.B) {
	var ops uint64
	for i := 0; i < b.N; i++ {
		atomic.AddUint64(&ops, 1)
	}
}

func benchmarkAddInt32(b *testing.B) {
	var ops int32 = 0
	for i := 0; i < b.N; i++ {
		atomic.AddInt32(&ops, 1)
	}
}
