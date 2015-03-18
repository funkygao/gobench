package main

import (
	"github.com/funkygao/gobench/util"
	"testing"
	"time"
)

func main() {
	b := testing.Benchmark(benchmarkNow)
	util.ShowBenchResult("time.Now", b)
	b = testing.Benchmark(benchmarkTimeSince)
	util.ShowBenchResult("time.Since", b)
	b = testing.Benchmark(benchmarkSleep)
	util.ShowBenchResult("time.Sleep", b)
	b = testing.Benchmark(benchmarkTimeUnixNano)
	util.ShowBenchResult("time.Now.UnixNano", b)
}

func benchmarkNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Now()
	}
}

func benchmarkSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Sleep(0)
	}
}

func benchmarkTimeSince(b *testing.B) {
	t1 := time.Now()
	for i := 0; i < b.N; i++ {
		time.Since(t1)
	}
}

func benchmarkTimeUnixNano(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		time.Now().UnixNano()
	}
}
