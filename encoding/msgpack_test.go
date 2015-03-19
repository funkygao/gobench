package encoding

import (
	"github.com/funkygao/golib/rand"
	"github.com/vmihailenco/msgpack"
	"testing"
)

func BenchmarkMsgpackMarshal(b *testing.B) {
	b.ReportAllocs()
	v := make(map[string]string)
	b.StopTimer()
	for i := 0; i < 10; i++ {
		v[rand.RandomString(5)] = rand.RandomString(7)
	}
	b.Logf("%+v", v)
	b.StartTimer()
	var data []byte
	for i := 0; i < b.N; i++ {
		data, _ = msgpack.Marshal(v)
	}
	b.SetBytes(int64(len(data)))
}

func BenchmarkMsgpackUnmarshal(b *testing.B) {
	b.ReportAllocs()
	v := make(map[string]string)
	b.StopTimer()
	for i := 0; i < 10; i++ {
		v[rand.RandomString(5)] = rand.RandomString(7)
	}
	data, _ := msgpack.Marshal(v)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		msgpack.Unmarshal(data)
	}
	b.SetBytes(int64(len(data)))
}
