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
BenchmarkInternalWithStack-8             1673311               685.7 ns/op          1536 B/op         11 allocs/op
BenchmarkInternalWithoutStack-8         144404904                8.213 ns/op           0 B/op          0 allocs/op
BenchmarkInternaEmptylWithStack-8        1567843               757.4 ns/op          1536 B/op         11 allocs/op
BenchmarkInternalEmptyWithoutStack-8    252769448                4.510 ns/op           0 B/op          0 allocs/op
BenchmarkNewEmptyWithStack-8             1701417               709.9 ns/op          1536 B/op         11 allocs/op
BenchmarkNewEmptylWithoutStack-8        266327073                4.640 ns/op           0 B/op          0 allocs/op
BenchmarkNewWithStack-8                  1658805               718.7 ns/op          1536 B/op         11 allocs/op
BenchmarkNewlWithoutStack-8             117173900                9.539 ns/op           0 B/op          0 allocs/op
```

![operations](./assets/operations.png)
![time operations](./assets/time_operations.png)

## Credit:
Based off work in:

- [stackus/errors](https://github.com/stackus/errors/tree/master)
- [pacman/errors](https://github.com/segmentfault/pacman/tree/main/errors)
