package main

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"unsafe"
)

type iface struct {
	tab  uintptr
	data unsafe.Pointer
}

func inpsectBytesBuffer() {
	var r io.Reader
	var w io.Writer
	var rw1 io.ReadWriter
	var rw2 io.ReadWriter

	buf := new(bytes.Buffer)
	rw1 = buf
	rw2 = buf
	w = buf

	// 可以将一个接口赋值给另一个接口，主要前者方法集包含目标接口方法集。
	// rw1 运行时iface结构中的 data 复制一份到 r 中的 data 字段
	r = rw1

	fmt.Printf("buf: %p\n", buf)

	irw1 := (*iface)(unsafe.Pointer(&rw1))
	fmt.Println("rw1 iface:", irw1)

	irw2 := (*iface)(unsafe.Pointer(&rw2))
	fmt.Println("rw2 iface:", irw2)

	iw := (*iface)(unsafe.Pointer(&w))
	fmt.Println("w iface:", iw)

	ir := (*iface)(unsafe.Pointer(&r))
	fmt.Println("r iface:", ir)
}

type SyntaxError struct {
	Num string
}

func (e SyntaxError) Error() string {
	return "parsing " + strconv.Quote(e.Num) + ": invalid syntax"
}

func inspectSyntaxError() {
	var err1, err2 error

	sErr := SyntaxError{Num: "x"}
	err1 = sErr
	err2 = sErr

	// err1 和 err2 运行时iface结构中的 data 分别都复制了sErr。
	// 所以通常我们都会选择指针对象来实现需要的接口类型
	iErr1 := (*iface)(unsafe.Pointer(&err1))
	fmt.Println("err1 iface:", iErr1)

	iErr2 := (*iface)(unsafe.Pointer(&err2))
	fmt.Println("err2 iface:", iErr2)

	// 修改 sErr 中的 Num 并不会影响 err1 和 err2
	sErr.Num = "y"
	fmt.Println("err1:", err1)
	fmt.Println("err2:", err2)

	// 接口仅在赋值时拷贝对象，后续复制接口时仅拷贝"iface.data"指针，底层对象是同一个
	func(err error) {
		iErr := (*iface)(unsafe.Pointer(&err))
		(*SyntaxError)(iErr.data).Num = "z"
	}(err1)
	fmt.Println("err1:", err1)
}

func main() {
	inpsectBytesBuffer()
	fmt.Println("==================")
	inspectSyntaxError()
}
