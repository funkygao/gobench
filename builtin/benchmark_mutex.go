package main

import (
	"github.com/funkygao/gobench/util"
	"sync"
	"testing"
)

func main() {
	b := testing.Benchmark(benchmarkMutexLock)
	util.ShowBenchResult("mutex.lock+unlock", b)
}

func benchmarkMutexLock(b *testing.B) {
	var lock sync.Mutex
	for i := 0; i < b.N; i++ {
		lock.Lock()
		lock.Unlock()
	}
}
