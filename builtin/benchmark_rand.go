package main

import (
	"github.com/funkygao/gobench/util"
	"math/rand"
	"testing"
)

func main() {
	b := testing.Benchmark(benchmarkRandInt31)
	util.ShowBenchResult("rand.Int31", b)
}

func benchmarkRandInt31(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		rand.Int31()
	}
}
