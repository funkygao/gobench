package main

import (
	"fmt"
	"unsafe"
)

type MyStruct struct {
	i int32
	j int64
}

func (this *MyStruct) Println() {
	fmt.Printf("i=%d, j=%d\n", this.i, this.j)
}

func main() {
	v := new(MyStruct)
	v.Println()

	// get v.i addr
	var i *int32 = (*int32)(unsafe.Pointer(v))
	*i = int32(8)
	v.Println()

	// try get v.j addr with failure because of padding/alignment
	var j *int64 = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + uintptr(unsafe.Sizeof(int32(0)))))
	*j = int64(763)
	v.Println()

	// get v.j addr by Alignof
	var k *int64 = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + uintptr(unsafe.Alignof(v.j))))
	*k = int64(763)
	v.Println()

	// get v.j addr by Offsetof
	var m *int64 = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + unsafe.Offsetof(v.j)))
	*m = 5699
	v.Println()

	// get v.j addr by pointer directly
	var n *int64 = (*int64)(unsafe.Pointer(&v.j))
	*n = 8911
	v.Println()

	fmt.Println(i, j, &v.j, k, m, n)
}
