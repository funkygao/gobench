package main

import (
	"github.com/funkygao/gobench/util"
	"testing"
)

func main() {
	b := testing.Benchmark(benchmarkDefer)
	util.ShowBenchResult("defer", b)
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
