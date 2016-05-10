# slice

Slice is a library for the [Go Programming Language][go]. It provides operations to work
with slices.

## Status

[![Build Status](https://img.shields.io/travis/raiqub/slice/master.svg?style=flat&label=linux%20build)](https://travis-ci.org/raiqub/slice)
[![AppVeyor Build](https://img.shields.io/appveyor/ci/skarllot/slice/master.svg?style=flat&label=windows%20build)](https://ci.appveyor.com/project/skarllot/slice)
[![GoDoc](https://godoc.org/github.com/raiqub/slice?status.svg)](http://godoc.org/github.com/raiqub/slice)

## Installation

To install raiqub/slice library run the following command:

```bash
go get gopkg.in/raiqub/slice.v1
```

To import this package, add the following line to your code:

```bash
import "gopkg.in/raiqub/slice.v1"
```

## Examples

```Go
var (
	SampleTextArray = []string{
		"Lorem", "ipsum", "dolor", "sit", "amet,", "consectetur", "adipiscing",
		"elit", "Sed", "tortor", "justo", "dui", "iaculis", "molestie",
		"Integer",
	}
)

// The out result will have all words that has more than 3 characters.
out := slice.String(SampleTextArray).Where(func(s string) bool {
	return len(s) > 3
})
```

```Go
var (
	SampleIntArray = []int{
		296, 112, 380, 243, 336,
		376, 664, 556, 162, 173,
		684, 503, 542, 215, 607,
		2, 132, 539, 646, 205,
	}
)

// The avg result will be the average of all slice elements (7373)
avg := slice.NInt(SampleIntArray).Average()
```

For reference and examples browse [library documentation][doc].

## Running tests

The tests can be run via provided `Makefile` script:

```bash
make test
```

## License

raiqub/slice is made available under the [Apache Version 2.0 License][license].


[go]: http://golang.org/
[doc]: http://godoc.org/github.com/raiqub/slice
[license]: http://www.apache.org/licenses/LICENSE-2.0