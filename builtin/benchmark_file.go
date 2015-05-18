package main

import (
	"github.com/funkygao/gobench/util"
	"math/rand"
	"os"
	"testing"
)

func main() {
	b := testing.Benchmark(benchmarkFileSeekSameOffset)
	util.ShowBenchResult("file.Seek same offset", b)
	b = testing.Benchmark(benchmarkFileSeekSameOffsetWithCache)
	util.ShowBenchResult("file.Seek same offset with cache", b)
	b = testing.Benchmark(benchmarkFileSeekRandOffset)
	util.ShowBenchResult("file.Seek rand offset", b)
}

func benchmarkFileSeekSameOffset(b *testing.B) {
	f, err := os.Open("benchmark_file.go")
	defer f.Close()
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		f.Seek(100, 0)
	}
}

func benchmarkFileSeekSameOffsetWithCache(b *testing.B) {
	f, err := os.Open("benchmark_file.go")
	defer f.Close()
	if err != nil {
		b.Fatal(err)
	}
	seeked := false
	for i := 0; i < b.N; i++ {
		if !seeked {
			f.Seek(100, 0)
			seeked = true
		}
	}
}

func benchmarkFileSeekRandOffset(b *testing.B) {
	f, err := os.Open("benchmark_file.go")
	defer f.Close()
	if err != nil {
		b.Fatal(err)
	}
	st, _ := f.Stat()
	for i := 0; i < b.N; i++ {
		f.Seek(rand.Int63()%st.Size(), 0)
	}
}
