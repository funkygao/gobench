// A func returning a slice/map or passing to another func, usually incurs a heap alloc.
//
// GC rate.
// Next GC is after we've allocated an extra amount of memory proportional to
// the amount already in use. The proportion is controlled by GOGC environment variable
// (100 by default). If GOGC=100 and we're using 4M, we'll GC again when we get to 8M.
//
// runtime.MemProfileRate to control heap profiling rate.

package main
