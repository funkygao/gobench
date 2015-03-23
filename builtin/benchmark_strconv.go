package main

import (
	"github.com/funkygao/gobench/util"
	"strconv"
	"testing"
)

func main() {
	b := testing.Benchmark(benchmarkStrconvAtoi)
	util.ShowBenchResult("Atoi", b)

}

func benchmarkStrconvAtoi(b *testing.B) {
	b.ReportAllocs()
	s := "1212"
	for i := 0; i < b.N; i++ {
		strconv.Atoi(s)
	}
}
