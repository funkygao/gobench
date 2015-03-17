package main

import (
	"github.com/funkygao/gobench/util"
	"testing"
)

func main() {
	b := testing.Benchmark(benchmarkSliceAppend)
	util.ShowBenchResult("slice.append", b)
}

func benchmarkSliceAppend(b *testing.B) {
	b.ReportAllocs()
	slice := make([]int, 0)
	for i := 0; i < b.N; i++ {
		slice = append(slice, i)
	}
}
