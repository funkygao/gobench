package main

import (
	"bytes"
	"fmt"
	"github.com/funkygao/gobench/util"
	"strings"
	"testing"
)

var (
	s1 = "hello "
	s2 = "world"
)

func main() {
	b := testing.Benchmark(BenchmarkStringContatWithPlus)
	util.ShowBenchResult("string a+b", b)
	b = testing.Benchmark(BenchmarkStringContatWithPlus1)
	util.ShowBenchResult("string a+=b", b)
	b = testing.Benchmark(BenchmarkStringContatWithFmt)
	util.ShowBenchResult("string fmt.Sprintf", b)
	b = testing.Benchmark(BenchmarkStringContatWithBytes)
	util.ShowBenchResult("string bytes.Buffer", b)
	b = testing.Benchmark(BenchmarkStringContatWithBytes1)
	util.ShowBenchResult("string bytes.Buffer1", b)
	b = testing.Benchmark(BenchmarkStringsJoin)
	util.ShowBenchResult("strings.join", b)
	b = testing.Benchmark(BenchmarkConcatWithCopy)
	util.ShowBenchResult("byte copy", b)
}

func BenchmarkStringContatWithPlus(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = s1 + s2
	}
	b.SetBytes(int64(len(s1) + len(s2)))
}

func BenchmarkStringContatWithPlus1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		x := s1
		x += s2
		_ = x
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
		_ = buf.String()
		buf.Reset()
	}
	b.SetBytes(int64(len(s1) + len(s2)))
}

func BenchmarkStringContatWithBytes1(b *testing.B) {
	b.ReportAllocs()
	buf := &bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		buf.Grow(len(s1) + len(s2))
		buf.WriteString(s1)
		buf.WriteString(s2)
		_ = buf.String()
		buf.Reset()
	}
	b.SetBytes(int64(len(s1) + len(s2)))
}

func BenchmarkStringsJoin(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = strings.Join([]string{s1, s2}, "")
	}
	b.SetBytes(int64(len(s1) + len(s2)))
}

func BenchmarkConcatWithCopy(b *testing.B) {
	b.ReportAllocs()
	bs := make([]byte, len(s1)+len(s2))
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(s1); j++ {
			bs[j] = s1[j]
		}
		for j := len(s1); j < len(s1)+len(s2); j++ {
			bs[j] = s2[j-len(s1)]
		}

		_ = string(bs)
	}
	b.SetBytes(int64(len(s1) + len(s2)))
}
