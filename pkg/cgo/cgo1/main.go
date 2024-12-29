// Cgo transforms the specified input Go source files into several output Go
// and C source files.

// _cgo_runtime_cgocall 对应 runtime.cgocall 函数
// https://github.com/golang/go/blob/go1.18.3/src/runtime/cgocall.go#L123
package main

//go:generate go tool cgo main.go

// int sum(int a, int b) { return a+b; }
import "C"

func main() {
	println(C.sum(1, 1))
}
