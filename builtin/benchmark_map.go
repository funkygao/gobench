package main

import (
	"github.com/funkygao/gobench/util"
	"github.com/funkygao/golib/rand"
	"testing"
)

func main() {
	b := testing.Benchmark(benchmarkMapSet)
	util.ShowBenchResult("map[x]=true", b)
	b = testing.Benchmark(benchmarkMapKeyStringLen10)
	util.ShowBenchResult("map[string len10] get", b)
	b = testing.Benchmark(benchmarkMapKeyStringLen100)
	util.ShowBenchResult("map[string len100] get", b)
}

func benchmarkMapSet(b *testing.B) {
	var x map[int]bool = make(map[int]bool)
	for i := 0; i < b.N; i++ {
		x[i] = true
	}
}

func benchmarkMapKeyStringLen10(b *testing.B) {
	const N = 10
	var x map[string]bool = make(map[string]bool)
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		x[rand.RandomString(N)] = true
	}

	b.StartTimer()
	key := rand.RandomString(N)
	for i := 0; i < b.N; i++ {
		_ = x[key]
	}
}

func benchmarkMapKeyStringLen100(b *testing.B) {
	const N = 100
	var x map[string]bool = make(map[string]bool)
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		x[rand.RandomString(N)] = true
	}

	b.StartTimer()
	key := rand.RandomString(N)
	for i := 0; i < b.N; i++ {
		_ = x[key]
	}
}
