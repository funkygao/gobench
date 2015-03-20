package hash_test

import (
	"github.com/funkygao/golib/hack"
	"hash/adler32"
	"hash/crc32"
	"hash/fnv"
	"strings"
	"testing"
)

func BenchmarkCrc32Of100B(b *testing.B) {
	b.ReportAllocs()
	key := strings.Repeat("s", 100)
	for i := 0; i < b.N; i++ {
		crc32.ChecksumIEEE([]byte(key))
	}
	b.SetBytes(100)
}

func BenchmarkAdler32Of100B(b *testing.B) {
	b.ReportAllocs()
	key := strings.Repeat("s", 100)
	f := func(k string) {
		//adler32.Checksum([]byte(k))
		adler32.Checksum(hack.Byte(k))
	}
	for i := 0; i < b.N; i++ {
		f(key)
	}
	b.SetBytes(100)
}

func BenchmarkFnv100B(b *testing.B) {
	b.ReportAllocs()
	key := strings.Repeat("s", 100)
	hasher := fnv.New32()
	f := func(k string) {
		hasher.Write(hack.Byte(k))
		hasher.Sum32()
	}

	for i := 0; i < b.N; i++ {
		f(key)
	}
	b.SetBytes(100)
}
