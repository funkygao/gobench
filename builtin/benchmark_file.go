package main

import (
	"github.com/funkygao/gobench/util"
	"io/ioutil"
	"os"
	"testing"
)

func main() {
	b := testing.Benchmark(benchmarkReadFile)
	util.ShowBenchResult("iouti.ReadFile", b)
}

func benchmarkReadFile(b *testing.B) {
	const F = "benchmark_file.go"

	for i := 0; i < b.N; i++ {
		ioutil.ReadFile(F)
	}

	f, _ := os.Open(F)
	st, _ := f.Stat()
	b.SetBytes(st.Size())
}
