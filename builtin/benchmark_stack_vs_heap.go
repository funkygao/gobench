package main

import (
	"github.com/funkygao/gobench/util"
	"testing"
)

func main() {
	b := testing.Benchmark(benchmarkStackAllocation)
	util.ShowBenchResult("stack alloc:", b)
	b = testing.Benchmark(benchmarkHeapAllocation)
	util.ShowBenchResult("heap  alloc:", b)
}

func stackalloc() [128]int64 {
	return [128]int64{}
}

func heapalloc() *[128]int64 {
	return &[128]int64{}
}

func benchmarkStackAllocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stackalloc()
	}
}

func benchmarkHeapAllocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		heapalloc()
	}
}
