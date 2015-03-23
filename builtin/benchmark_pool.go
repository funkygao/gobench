package main

import (
	"github.com/funkygao/gobench/util"
	"sync"
	"testing"
)

func main() {
	b := testing.Benchmark(benchmarkSyncPoolGetThenPut)
	util.ShowBenchResult("sync.pool get & put []byte", b)
	b = testing.Benchmark(benchmarkSyncPoolGetThenPutWithPointer)
	util.ShowBenchResult("sync.pool get & put *[]byte", b)
}

func benchmarkSyncPoolGetThenPut(b *testing.B) {
	b.ReportAllocs()
	var p = sync.Pool{New: func() interface{} {
		return make([]byte, 1024)
	}}
	for i := 0; i < b.N; i++ {
		x := p.Get()
		p.Put(x)
	}
}

func benchmarkSyncPoolGetThenPutWithPointer(b *testing.B) {
	b.ReportAllocs()
	var p = sync.Pool{New: func() interface{} {
		b := make([]byte, 1024)
		return &b
	}}
	for i := 0; i < b.N; i++ {
		x := p.Get().(*[]byte)
		p.Put(x)
	}
}
