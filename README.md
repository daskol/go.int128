# Go.Int128

*trial implemntation of 128 bit arithmetics in Go*

## Overview

Now only summation of 128 bit integers are implemented. The benchmark below
shows supremacy of the custom implentation.

```
goos: linux
goarch: amd64
pkg: github.com/daskol/go.int128
BenchmarkAdd/Simple-4               2000000000           0.37 ns/op
BenchmarkBigAdd/Simple-4            100000000           17.1  ns/op
BenchmarkInt128Add/Simple-4         500000000            3.62 ns/op
PASS
ok      github.com/daskol/go.int128 4.697s
```
