package encoding

import (
	"github.com/funkygao/fae/servant/gen-go/fun/rpc"
	"github.com/funkygao/thrift/lib/go/thrift"
	"testing"
)

func BenchmarkThriftSerialize(b *testing.B) {
	b.ReportAllocs()

	transport := thrift.NewTMemoryBuffer()
	protocol := thrift.NewTBinaryProtocolFactoryDefault()
	//iprot := protocol.GetProtocol(transport)
	oprot := protocol.GetProtocol(transport)
	mysqlResult := rpc.NewMysqlResult()
	mysqlResult.Cols = make([]string, 0)
	mysqlResult.Rows = make([][]string, 0)
	colsN, rowsN := 5, 10
	for i := 0; i < colsN; i++ {
		mysqlResult.Cols = append(mysqlResult.Cols, "username")
	}
	for i := 0; i < rowsN; i++ {
		row := make([]string, 0)
		for j := 0; j < colsN; j++ {
			row = append(row, "beijing, los angels")
		}

		mysqlResult.Rows = append(mysqlResult.Rows, row)
	}
	//b.Logf("%s", mysqlResult.String())

	for i := 0; i < b.N; i++ {
		mysqlResult.Write(oprot)
	}

	transport.Close()
	b.SetBytes(int64(len(mysqlResult.String())))
}
