package main

import (
	"github.com/funkygao/gobench/util"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func main() {
	b := testing.Benchmark(benchmarkIoReadFull)
	util.ShowBenchResult("io.ReadFull", b)
	b = testing.Benchmark(benchmarkReadFile)
	util.ShowBenchResult("iouti.ReadFile", b)
}

func benchmarkReadFile(b *testing.B) {
	const F = "benchmark_io.go"

	for i := 0; i < b.N; i++ {
		ioutil.ReadFile(F)
	}

	f, _ := os.Open(F)
	st, _ := f.Stat()
	b.SetBytes(st.Size())
}

func benchmarkIoReadFull(b *testing.B) {
	r := stringReader{}
	buf := make([]byte, 1000)
	for i := 0; i < b.N; i++ {
		io.ReadFull(r, buf)
	}
}

type stringReader struct {
}

func (this stringReader) Read(p []byte) (n int, err error) {
	p = []byte{4, 5, 99, 10, 111, 2, 4, 4, 4, 4, 4, 4, 4, 4}
	return len(p), nil
}
