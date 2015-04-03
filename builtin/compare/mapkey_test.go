package main

import (
	"testing"
)

func BenchmarkMapWithStringKey(b *testing.B) {
	s, n := "reasonably-long-but-present-unique-identifier",
		"non-present-unique-identifier"
	m := map[string]struct{}{s: struct{}{}}
	for i := 0; i < b.N; i++ {
		_, _ = m[s]
		_, _ = m[n]
	}
}

func BenchmarkMapWithStructKey(b *testing.B) {
	type key struct{ a, b int }
	k, n := key{0, 1}, key{1, 2}
	m := map[key]struct{}{k: struct{}{}}
	for i := 0; i < b.N; i++ {
		_, _ = m[k]
		_, _ = m[n]
	}
}
