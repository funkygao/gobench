package main

import (
	"github.com/funkygao/gobench/util"
	"sync"
	"testing"
)

func main() {
	b := testing.Benchmark(benchmarkMutexLockThenUnlock)
	util.ShowBenchResult("mutex.lock+unlock", b)
}

func benchmarkMutexLockThenUnlock(b *testing.B) {
	var lock sync.Mutex
	for i := 0; i < b.N; i++ {
		lock.Lock()
		lock.Unlock()
	}
}
