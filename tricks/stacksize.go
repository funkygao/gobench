// Test how big can the stack of a go routine be.
// runtime/proc.go: if 64OS maxstacksize=1GB else maxstacksize=250MB
//
// Allocate a bigger segment and move the stack:
// oldsize = gp->stack.hi - gp->stack.lo;
// newsize = oldsize * 2;
package main

func main() {
	main() // will panic: runtime: goroutine stack exceeds 1000000000-byte limit
}
