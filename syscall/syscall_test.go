package syscall

import (
	"os"
	"testing"
)

// 130ns/op
func BenchmarkSyscallOverhead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		os.Getpid()
	}
}
