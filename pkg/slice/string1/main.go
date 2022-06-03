// 字符串就是一个只读的字节切片
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 字符串运行时数据结构如下：
// type string struct {
// 	ptr unsafe.Pointer
// 	len int
// }

// 字符串作为参数传递时，拷贝的是 "string header"（大小为 2 word，64位机器是16字节）
func toUpper(s string) string {
	r := make([]byte, len(s))
	for i, b := range []byte(s) {
		if 'a' <= b && b <= 'z' {
			r[i] = b + 'A' - 'a'
		} else {
			r[i] = b
		}
	}
	// 不管是 []byte(s) 还是 string(r) 都会发生底层字节数组的拷贝
	return string(r)
}

// b2s 将字节切片转换为字符串，但不会产生内存分配
func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// s2b 将字符串转换为字节切片，但不会产生内存分配
func s2b(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return b
}

func main() {
	s1 := "hello world!"
	fmt.Println(toUpper(s1))

	s2 := s1[:5]
	fmt.Println(s2)

	// 截取部分字符串是非常高效的，底层共享同一个字节数组
	h1 := (*reflect.StringHeader)(unsafe.Pointer(&s1))
	h2 := (*reflect.StringHeader)(unsafe.Pointer(&s2))
	fmt.Printf("s1: %#x s2: %#x", h1.Data, h2.Data)
}
