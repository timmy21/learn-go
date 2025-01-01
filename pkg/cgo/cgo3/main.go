package main

/*
#cgo CFLAGS: -I${SRCDIR}/ctestlib
#cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/ctestlib
#cgo LDFLAGS: -L${SRCDIR}/ctestlib
#cgo LDFLAGS: -ltest

#include <stdlib.h>
#include "test.h"

char *foo = "hellofoo";
*/
import "C"

import (
	"fmt"
	"time"
	"unsafe"
)

func Random() int {
	return int(C.random()) // C.long -> int
}

func Seed(i int) {
	C.srandom(C.uint(i)) // int -> C.uint
}

func main() {
	Seed(int(time.Now().Unix()))
	fmt.Println(Random())

	s := "hello cstr"
	// 通过C.CString在C内部分配的内存，Go中GC是无法感知到的，
	// 因此要记着在使用后手动释放
	cs := C.CString(s)
	C.print_string(cs)
	C.free(unsafe.Pointer(cs))

	// C.GoString在Go世界重新分配一块内存对象，并复制了C的字符串的信息
	// 后续这个位于Go世界的新string类型对象和其他Go对象一样接受GC的管理
	fmt.Printf("%T\n", C.GoString(C.foo))
}
