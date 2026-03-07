# pathtype

📂 [`path`](https://pkg.go.dev/path) & [`filepath`](https://pkg.go.dev/path/filepath) functions exposed as methods on a path newtype

<table align=center><td>

```go
x := p.Dir().Join(q.Dir().Dir().Base())
```

</table>

## Installation

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=FFFFFF)

```sh
go get go.jcbhmr.com/pathtype
```

## Usage

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=FFFFFF)

This module offers two primary types: [`pathtype.Path`](https://pkg.go.dev/go.jcbhmr.com/pathtype#Path) and [`filepath.FilePath`](https://pkg.go.dev/go.jcbhmr.com/pathtype/filepath#FilePath). These are newtype wrappers around primitive strings to improve the ergonomics of reading nested path operation functions by changing an inside-out right-to-left nested parenthesis mess into a chain of methods that reads from left to right.

```go
// Before
x := filepath.Join(filepath.Dir(p), filepath.Base(filepath.Dir(filepath.Dir(q))))

// After
x := p.Dir().Join(q.Dir().Dir().Base())
xString := string(x) // if you want a string primitive
```

You are encouraged to use these types and their methods as a replacement for the standard library [`path`](https://pkg.go.dev/path) and [`filepath`](https://pkg.go.dev/path/filepath) packages.

[📚 Check out the documentation for more examples & API details](https://pkg.go.dev/go.jcbhmr.com/pathtype)

## Development

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=FFFFFF)

You'll need Go to develop this module. You can run some basic tests and make sure everything compiles using `go test -v ./...`.

**Why is `pathtype` named after `path` with a `*type` suffix but `filepath` is not named `filepathtype`?** \
Because `pathtype` is the module's name and I didn't want to create a redundant `go.jcbhmr.com/pathtype/path` subpackage and have an empty root package.
