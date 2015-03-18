package main

import (
	"github.com/funkygao/gobench/util"
	"testing"
)

func main() {
	b := testing.Benchmark(benchmarkTypeAssert)
	util.ShowBenchResult("type assert", b)
}

func benchmarkTypeAssert(b *testing.B) {
	var a interface{} = 45
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, ok := a.(int); ok {

		}
	}
}
