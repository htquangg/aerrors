# AErrors

This is a fork of @stackus's [errors](https://github.com/stackus/errors/tree/master). I'll be maintaining this version for my own use.

## Installation

    go get -u github.com/htquangg/aerrors

## Prerequisites

Go 1.22

## Benchmarks

```shell
▶ go test ./... -test.run=NONE -test.bench=. -test.benchmem
goos: darwin
goarch: arm64
pkg: github.com/htquangg/aerrors
BenchmarkNewWithoutStack-8      347110820                3.305 ns/op           0 B/op          0 allocs/op
BenchmarkNewWithStack-8           433795              2682 ns/op            1656 B/op         13 allocs/op
BenchmarkRawWithoutStack-8      358170732                3.383 ns/op           0 B/op          0 allocs/op
BenchmarkRawWithStack-8           418533              2685 ns/op            1656 B/op         13 allocs/op
```

## Credit:
Based off work in:

- [stackus/errors](https://github.com/stackus/errors/tree/master)
