package main

import (
	"github.com/funkygao/gobench/util"
	"reflect"
	"testing"
)

func main() {
	b := testing.Benchmark(benchmarkDeepEqual)
	util.ShowBenchResult("reflect.DeepEqual", b)
	b = testing.Benchmark(benchmarkValueOf)
	util.ShowBenchResult("reflect.ValueOf", b)
	b = testing.Benchmark(benchmarkTypeOf)
	util.ShowBenchResult("reflect.TypeOf", b)
}

func benchmarkDeepEqual(b *testing.B) {
	x, y := []string{"a", "b"}, []string{"c", "d", "e"}
	for i := 0; i < b.N; i++ {
		reflect.DeepEqual(x, y)
	}
	b.SetBytes(5)
}

func benchmarkValueOf(b *testing.B) {
	var foo = map[string]interface{}{"ff": 3.4}
	for i := 0; i < b.N; i++ {
		reflect.ValueOf(foo)
	}
}

func benchmarkTypeOf(b *testing.B) {
	var foo = map[string]interface{}{"ff": 3.4}
	for i := 0; i < b.N; i++ {
		reflect.TypeOf(foo)
	}
}
