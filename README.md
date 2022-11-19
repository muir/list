# list - golang generics for list manipulation

[![GoDoc](https://godoc.org/github.com/muir/list?status.png)](https://pkg.go.dev/github.com/muir/list)
![unit tests](https://github.com/muir/list/actions/workflows/go.yml/badge.svg)
[![report card](https://goreportcard.com/badge/github.com/muir/list)](https://goreportcard.com/report/github.com/muir/list)
[![codecov](https://codecov.io/gh/muir/list/branch/main/graph/badge.svg)](https://codecov.io/gh/muir/list)

Install:

	go get github.com/muir/list

---

This package is a collection of generic functions to manipulate slices.


## Slice/Replace

```go
func Replace[E any](dest []E, start int, replace ...E) []E 
func ReplaceBeyond[E any](dest []E, start int, replace ...E) []E
func Splice[E any](dest []E, start int, end int, replace ...E) []E
func SpliceBeyond[E any](dest []E, start int, end int, replace ...E) []E
```

`Splice` is the general function of replacing one part of a list/vector/slice with
new elements, growing or shrinking the list as needed.

`Replace` is splicing such that the replacements equals the size replaced.

The `Beyond` methods allow the replacments to be beyond the len (and cap) of the
original slice.  The slice will be extended (and filled) as needed.

## Code stability

This is brand new, but has 100% test coverage and is unlikely to be modifed except
to add new functions.

## Contributions

I would be happy to collect other list functions in this repo.  Open a pull request.
Include tests that maintain the current 100% coverage.

