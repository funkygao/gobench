package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i uint32 = 0x12345678
	x := (*[4]byte)(unsafe.Pointer(&i))
	fmt.Printf("0x%x in memory: 0x", i)
	for _, xn := range *x {
		fmt.Printf("%x", xn)
	}
	fmt.Println()
}
