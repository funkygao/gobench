package main

import (
	"github.com/funkygao/gobench/util"
	"testing"
	"time"
)

func main() {
	b := testing.Benchmark(benchmarkNow)
	util.ShowBenchResult("time.Now", b)
	b = testing.Benchmark(benchmarkSleep)
	util.ShowBenchResult("time.Sleep", b)
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
