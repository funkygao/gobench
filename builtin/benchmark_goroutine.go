package main

import (
	"github.com/funkygao/gobench/util"
	"testing"
)

func main() {
	b := testing.Benchmark(benchmarkGoroutine)
	util.ShowBenchResult("go", b)
}

func benchmarkGoroutine(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		go func() {
		}()
	}
}
