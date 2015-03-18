package main

import (
	"github.com/funkygao/gobench/util"
	"testing"
)

func main() {
	b := testing.Benchmark(benchmarkMapSet)
	util.ShowBenchResult("map[x]=true", b)
}

func benchmarkMapSet(b *testing.B) {
	var x map[int]bool = make(map[int]bool)
	for i := 0; i < b.N; i++ {
		x[i] = true
	}
}
