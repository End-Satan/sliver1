# native [![Go Reference](https://pkg.go.dev/badge/github.com/josharian/native.svg)](https://pkg.go.dev/github.com/josharian/native)

Package native provides easy access to native byte order.

`go get github.com/josharian/native`

Usage: Use `native.Endian` where you need the native binary.ByteOrder.

Please think twice before using this package.
It can break program portability.
Native byte order is usually not the right answer.

