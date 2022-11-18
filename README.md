[![linters](https://github.com/SkippyZA/go-slices/actions/workflows/lint.yml/badge.svg)](https://github.com/SkippyZA/go-slices/actions/workflows/lint.yml)
[![test](https://github.com/SkippyZA/go-slices/actions/workflows/test.yml/badge.svg)](https://github.com/SkippyZA/go-slices/actions/workflows/test.yml) 
[![codecov](https://codecov.io/gh/SkippyZA/go-slices/branch/master/graph/badge.svg?token=OLBD9XQKZX)](https://codecov.io/gh/SkippyZA/go-slices)
![release](https://img.shields.io/github/v/release/skippyza/go-slices?include_prereleases)

# go-slices

The goal in mind for this `go-slices` module is to provide basic generics to cover all the common operations a developer will typically perform on a slice.

## Installation

```bash
go get github.com/skippyza/go-slices
```

## Examples

Filter even numbers
```golang
slices.Filter([]int{1, 2, 3, 4, 5}, func(i int) bool { return i%2 == 0 })
// [2 4]
```

Double all the numbers
```golang
double := func(i int) int { return i*2 }
slices.Map([]int{8, 14, 23}, double)
// [16 28 46]
```

Sum all the numbers in a slice
```golang
sum := func(a, b int) int { return a+b }
slices.Reduce([]int{5, 1, 2, 2}, sum, 0)
// 10
```

## Usage
See [DOCUMENTATION](https://pkg.go.dev/github.com/skippyza/go-slices) for more info.
