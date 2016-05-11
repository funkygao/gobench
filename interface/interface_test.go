package iface

import (
	"testing"
)

type I interface {
	Test(int)
}

type S struct {
	x int
}

func (this *S) Test(x int) {
	this.x = x
}

var data = &S{2}

func callStruct(d *S) {
	d.Test(10)
}

func callInteface(d I) {
	d.Test(10)
}

func BenchmarkStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		callStruct(data)
	}
}

// 7 times slower than benchmark struct
// because the dynamic binding has extra overhead
// go build -gcflags "-m" to see why
func BenchmarkInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		callInteface(data)
	}
}
