// Simulate a program that creates a lot of garbage.
// Once a second it creates a byte array of 5-10MB
package main

import (
	"github.com/funkygao/golib/gcvis"
	"github.com/funkygao/golib/gofmt"
	"github.com/funkygao/golib/terminal"
	"log"
	"math/rand"
	"runtime"
	"time"
)

func makeBuffer() []byte {
	return make([]byte, rand.Intn(5000000)+5000000)
}

func main() {
	gcvis.Launch(":8990", "gc", 2)
	log.Println("gcvis on :8990")

	pool := make([][]byte, 20)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	th := terminal.TerminalHeight()
	if th == 0 {
		th = 10
	}

	var m runtime.MemStats

	loops := 0
	for _ = range ticker.C {
		b := makeBuffer()
		pool[rand.Intn(len(pool))] = b // old ref will be GC'ed
		if loops%th == 0 {
			log.Printf("%10s %10s %10s %10s %10s",
				"HeapSys", "make", "HeapAlloc", "HeapIdle", "HeapReleased")
		}
		loops += 1

		bytes := 0
		for i := 0; i < len(pool); i++ {
			if pool[i] != nil {
				bytes += len(pool[i])
			}
		}

		runtime.ReadMemStats(&m)
		log.Printf("%10s %10s %10s %10s %10s",
			gofmt.ByteSize(m.HeapSys),
			gofmt.ByteSize(bytes),
			gofmt.ByteSize(m.HeapAlloc),
			gofmt.ByteSize(m.HeapIdle),
			gofmt.ByteSize(m.HeapReleased))
	}
}
