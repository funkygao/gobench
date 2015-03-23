package main

import (
	r "crypto/rand"
	"github.com/funkygao/gobench/util"
	"math/rand"
	"testing"
)

func main() {
	b := testing.Benchmark(benchmarkRandInt31)
	util.ShowBenchResult("rand.Int31", b)
	b = testing.Benchmark(benchmarkRandRead10)
	util.ShowBenchResult("rand.Read10", b)
	b = testing.Benchmark(benchmarkRandRead100)
	util.ShowBenchResult("rand.Read100", b)
}

func benchmarkRandInt31(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		rand.Int31()
	}
}

func benchmarkRandRead10(b *testing.B) {
	b.ReportAllocs()
	bytes := make([]byte, 10)
	for i := 0; i < b.N; i++ {
		r.Read(bytes)
	}
}

func benchmarkRandRead100(b *testing.B) {
	b.ReportAllocs()
	bytes := make([]byte, 100)
	for i := 0; i < b.N; i++ {
		r.Read(bytes)
	}
}
