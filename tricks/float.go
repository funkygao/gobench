package main

import (
	"fmt"
	"unsafe"
	"strconv"
)

func main() {
	a := float32(0.1)
	fmt.Printf("float32(0.1): %032s\n", strconv.FormatUint((uint64)(*(*uint32)(unsafe.Pointer(&a))),2))
	a = 0.2
	fmt.Printf("float32(0.2): %032s\n", strconv.FormatUint((uint64)(*(*uint32)(unsafe.Pointer(&a))),2))
	a = 0.3
	fmt.Printf("float32(0.3): %032s\n", strconv.FormatUint((uint64)(*(*uint32)(unsafe.Pointer(&a))),2))
	
	b := float64(0.1)
	fmt.Printf("float64(0.1): %064s\n", strconv.FormatUint((*(*uint64)(unsafe.Pointer(&b))),2))
	b = 0.2
	fmt.Printf("float64(0.2): %064s\n", strconv.FormatUint((*(*uint64)(unsafe.Pointer(&b))),2))
	b = 0.3
	fmt.Printf("float64(0.3): %064s\n", strconv.FormatUint((*(*uint64)(unsafe.Pointer(&b))),2))
}
