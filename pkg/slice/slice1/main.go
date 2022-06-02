package main

import (
	"fmt"
	"unsafe"
)

// https://github.com/golang/go/blob/master/src/runtime/slice.go#L15
// slice 的数据结构定义如下:
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

// append 元素时，如果 slice 容量不足以容纳新增元素时，会发生扩容。扩容规则如下:
//   1. 小容量(cap < 256)时，新容量 = 2 * 旧容量
//   2. 超过256时：Transition from growing 2x for small slices to growing 1.25x for large slices
// 不同Go版本扩容规则可能会发生变化，具体请查看：https://github.com/golang/go/blob/master/src/runtime/slice.go#L178
func concat(a, b []int) []int {
	// slice 做出参数传递时，拷贝的是 slice 结构体（例如：64位机器是24字节），而不会拷贝底层的数组。
	return append(a, b...)
}

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	return fmt.Sprintf("{X:%d Y:%d}", p.X, p.Y)
}

// slice 低成本拷贝，使用方式非常的灵活，所以使用 Go 编程时会大量使用。而数组仅在一些特定场景使用。
// 关于 slice 的各种操作可以查看：https://github.com/golang/go/wiki/SliceTricks
func delete(a []*Point, i int) []*Point {
	copy(a[i:], a[i+1:])
	a[len(a)-1] = nil // or the zero value of T，释放内存
	return a[:len(a)-1]
}

func main() {
	// make([]T, length, capacity)
	s1 := make([]int, 0, 6)

	s2 := concat(s1, []int{1, 2, 3})
	fmt.Println(s2)           // Output: [1 2 3]
	fmt.Println(cap(s2))      // Output: 6
	fmt.Println(s1)           // Output: []
	fmt.Println(s1[:cap(s1)]) // Output: [1 2 3 0 0 0]

	fmt.Println("==================")
	s3 := concat(s2, []int{4, 5, 6, 7})
	fmt.Println(s3)           // Output: [1 2 3 4 5 6 7]
	fmt.Println(cap(s3))      // Output: 12
	fmt.Println(s1)           // Output: []
	fmt.Println(s1[:cap(s1)]) // Output: [1 2 3 0 0 0]

	fmt.Println("==================")
	points := []*Point{
		{X: 0, Y: 1},
		{X: 1, Y: 2},
		{X: 2, Y: 3},
	}
	points = delete(points, 1)
	fmt.Println(points)
}
