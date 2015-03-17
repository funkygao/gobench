package main

import (
	"github.com/funkygao/gobench/util"
	"regexp"
	"strings"
	"testing"
)

func main() {
	b := testing.Benchmark(BenchmarkStringConstains)
	util.ShowBenchResult("strings.Contains", b)
	b = testing.Benchmark(BenchmarkStringConcat)
	util.ShowBenchResult("builtin string +", b)
	b = testing.Benchmark(BenchmarkStringLen)
	util.ShowBenchResult("len of string", b)
	b = testing.Benchmark(BenchmarkRegexpMatch)
	util.ShowBenchResult("regexp.NotCompiled", b)
	b = testing.Benchmark(BenchmarkRegexpMatchCompiled)
	util.ShowBenchResult("regexp.Compiled", b)
}

func BenchmarkStringConstains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Contains("we are all here", "all")
	}
	b.SetBytes(int64(len("we are all here")))
}

func BenchmarkStringConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = "home" + ":" + "nice"
	}
	b.SetBytes(int64(len("home" + ":" + "nice")))
}

func BenchmarkStringLen(b *testing.B) {
	s := "abcdefg"
	for i := 0; i < b.N; i++ {
		_ = len(s)
	}
	b.SetBytes(int64(len(s)))
}

func BenchmarkRegexpMatch(b *testing.B) {
	pattern := "child \\d+ started"
	line := "adfasdf  asdfas dfasdf child 12 started with asdfasf"
	for i := 0; i < b.N; i++ {
		regexp.MatchString(pattern, line)
	}
	b.SetBytes(int64(len(line)))
}

func BenchmarkRegexpMatchCompiled(b *testing.B) {
	pattern := regexp.MustCompile("child \\d+ started")
	line := "adfasdf  asdfas dfasdf child 12 started with asdfasf"
	for i := 0; i < b.N; i++ {
		pattern.MatchString(line)
	}
	b.SetBytes(int64(len(line)))
}
