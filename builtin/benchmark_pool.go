package main

import (
	"github.com/funkygao/gobench/util"
	"sync"
	"testing"
)

func main() {
	b := testing.Benchmark(benchmarkSyncPoolGetThenPut)
	util.ShowBenchResult("sync.pool get & put", b)
}

func benchmarkSyncPoolGetThenPut(b *testing.B) {
	b.ReportAllocs()
	var p = sync.Pool{New: func() interface{} {
		return make([]byte, 100)
	}}
	for i := 0; i < b.N; i++ {
		x := p.Get()
		p.Put(x)
	}
}
