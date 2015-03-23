package main

import (
	"bytes"
	"fmt"
	"github.com/funkygao/gobench/util"
	"testing"
)

var (
	s1 = "hello "
	s2 = "world"
)

func main() {
	b := testing.Benchmark(BenchmarkStringContatWithPlus)
	util.ShowBenchResult("string a+b", b)
	b = testing.Benchmark(BenchmarkStringContatWithFmt)
	util.ShowBenchResult("string fmt.Sprintf", b)
	b = testing.Benchmark(BenchmarkStringContatWithBytes)
	util.ShowBenchResult("string bytes.Buffer", b)
}

func BenchmarkStringContatWithPlus(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = s1 + s2
	}
	b.SetBytes(int64(len(s1) + len(s2)))
}

func BenchmarkStringContatWithFmt(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s", s1, s2)
	}
	b.SetBytes(int64(len(s1) + len(s2)))
}

func BenchmarkStringContatWithBytes(b *testing.B) {
	b.ReportAllocs()
	buf := &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		buf.WriteString(s1)
		buf.WriteString(s2)
		buf.Reset()
	}
	b.SetBytes(int64(len(s1) + len(s2)))
}
