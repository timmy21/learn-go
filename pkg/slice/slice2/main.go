package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 在运行时切片可以由 reflect.SliceHeader 结构体表示
// type SliceHeader struct {
// 	Data uintptr
// 	Len  int
// 	Cap  int
// }

func main() {
	// 切片数据结构中的 Data 存储的是 切片第一个元素在底层数组中的地址
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("arr: %v\n", arr)

	s1 := arr[0:6]
	h1 := (*reflect.SliceHeader)(unsafe.Pointer(&s1))
	// h1.Data = &arr[0]
	fmt.Printf("arr: %p s1: %#x len: %d cap: %d\n", &arr, h1.Data, h1.Len, h1.Cap)
	fmt.Printf("s1: %v\n", s1)

	s2 := arr[2:7]
	h2 := (*reflect.SliceHeader)(unsafe.Pointer(&s2))
	// h2.Data = &arr[2]
	fmt.Printf("arr: %p s2: %#x len: %d cap: %d\n", &arr, h2.Data, h2.Len, h2.Cap)
	fmt.Printf("s2: %v\n", s2)

	s3 := s2[1:6]
	h3 := (*reflect.SliceHeader)(unsafe.Pointer(&s3))
	// h3.Data = &arr[3]
	fmt.Printf("arr: %p s3: %#x len: %d cap: %d\n", &arr, h3.Data, h3.Len, h3.Cap)
	fmt.Printf("s3: %v\n", s3)
}
