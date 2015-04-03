package main

import (
	"github.com/funkygao/gobench/util"
	"testing"
)

func main() {
	b := testing.Benchmark(benchmarkMapVsStructOfMap)
	util.ShowBenchResult("MapVsStruct: map", b)
	b = testing.Benchmark(benchmarkMapVsStructOfStruct)
	util.ShowBenchResult("MapVsStruct: struct", b)
	b = testing.Benchmark(benchmarkMapWithStringKey)
	util.ShowBenchResult("map with key: string", b)
	b = testing.Benchmark(benchmarkMapWithStructKey)
	util.ShowBenchResult("map with key: struct", b)
}

func benchmarkMapVsStructOfMap(b *testing.B) {
	m := map[int]int{0: 0, 1: 1}
	for i := 0; i < b.N; i++ {
		_ = m[0] + m[1]
	}
}

func benchmarkMapVsStructOfStruct(b *testing.B) {
	m := struct{ a, b int }{0, 1}
	for i := 0; i < b.N; i++ {
		_ = m.a + m.b
	}
}

func benchmarkMapWithStringKey(b *testing.B) {
	s, n := "reasonably-long-but-present-unique-identifier",
		"non-present-unique-identifier"
	m := map[string]struct{}{s: struct{}{}}
	for i := 0; i < b.N; i++ {
		_, _ = m[s]
		_, _ = m[n]
	}
}

func benchmarkMapWithStructKey(b *testing.B) {
	type key struct{ a, b int }
	k, n := key{0, 1}, key{1, 2}
	m := map[key]struct{}{k: struct{}{}}
	for i := 0; i < b.N; i++ {
		_, _ = m[k]
		_, _ = m[n]
	}
}
