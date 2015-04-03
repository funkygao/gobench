#!/bin/sh
go test -c
./compare.test -test.bench="MapWithStringKey" -test.cpuprofile="MapWithStringKey.pprof"
./compare.test -test.bench="MapWithStructKey" -test.cpuprofile="MapWithStructKey.pprof"
go tool pprof -text compare.test MapWithStringKey.pprof > MapWithStringKey.out
go tool pprof -text compare.test MapWithStructKey.pprof > MapWithStructKey.out
rm -f *.pprof compare.test
