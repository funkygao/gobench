package util

import (
	"fmt"
	"testing"
)

func ShowBenchResult(name string, b testing.BenchmarkResult) {
	fmt.Printf("%s %s %s\n", name, b.String(), b.MemString())
}
