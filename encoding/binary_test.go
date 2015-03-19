package encoding

import (
	"encoding/binary"
	"testing"
)

func BenchmarkBinaryBigEndianPutUint16(b *testing.B) {
	b.ReportAllocs()
	v := make([]byte, 2)
	for i := 0; i < b.N; i++ {
		binary.BigEndian.PutUint16(v, 12)
	}
	b.SetBytes(int64(len(v)))
}

func BenchmarkBinaryBigEndianPutUint64(b *testing.B) {
	b.ReportAllocs()
	v := make([]byte, 8)
	for i := 0; i < b.N; i++ {
		binary.BigEndian.PutUint64(v, 12)
	}
	b.SetBytes(int64(len(v)))
}
