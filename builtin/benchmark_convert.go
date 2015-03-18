package main

import (
	"github.com/funkygao/gobench/util"
	"strings"
	"testing"
)

func main() {
	b := testing.Benchmark(benchmarkByteSliceConvertString)
	util.ShowBenchResult("[]byte(s.1000)", b)
	b = testing.Benchmark(benchmarkStringConvertByteSlice)
	util.ShowBenchResult("string(b.1000)", b)
}

func benchmarkByteSliceConvertString(b *testing.B) {
	b.ReportAllocs()
	s := strings.Repeat("h", 1000)
	for i := 0; i < b.N; i++ {
		_ = []byte(s) // will lead to mem copy and alloc
	}
}

// 1338 ns/op
func benchmarkStringConvertByteSlice(b *testing.B) {
	const N = 1000
	b.ReportAllocs()
	ba := make([]byte, N) // lower N will show better performance
	for i := 0; i < b.N; i++ {
		_ = string(ba) // its costly
	}
}
