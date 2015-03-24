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
	var i *int32 = (*int32)(unsafe.Pointer(v))
	*i = int32(8)
	var j *int64 = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + uintptr(unsafe.Sizeof(int32(0)))))
	*j = int64(763)
	v.Println()
	var k *int64 = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + uintptr(2*4)))
	*k = int64(763)
	v.Println()
	fmt.Println(i, j, &v.j, k)

	// j doesn't work while k works, because of golang padding
}