#!/bin/sh
echo ========================
echo golang builtin benchmark
echo ========================
go test -benchmem -cpu 1,2,4 -test.bench=".*" fmt
go test -benchmem -cpu 1,2,4 -test.bench=".*" bytes
go test -benchmem -cpu 1,2,4 -test.bench=".*" strconv
go test -benchmem -cpu 1,2,4 -test.bench=".*" regexp
go test -benchmem -cpu 1,2,4 -test.bench=".*" strings

echo
echo

for f in `ls benchm*.go`
do
    echo $f
    echo ========================
    go run $f
    echo
done
