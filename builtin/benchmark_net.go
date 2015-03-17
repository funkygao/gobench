package main

import (
	"github.com/funkygao/gobench/util"
	"net"
	"testing"
)

func main() {
	b := testing.Benchmark(benchmarkDialAndClose)
	util.ShowBenchResult("net.Dial+Close", b)
}

func benchmarkDialAndClose(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, e := net.Dial("tcp", "localhost:80")
		if e != nil {
			panic(e)
		}
		conn.Close()
	}
}
