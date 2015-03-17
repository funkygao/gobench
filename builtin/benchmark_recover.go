package main

import (
	"github.com/funkygao/gobench/util"
	"testing"
)

func main() {
	b := testing.Benchmark(benchmarkRecover)
	util.ShowBenchResult("recover", b)
}

func benchmarkRecover(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		recover()
	}
}
