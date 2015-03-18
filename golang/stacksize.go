// Test how big can the stack of a go routine be.
package main

func main() {
	main() // will panic: runtime: goroutine stack exceeds 1000000000-byte limit
}
