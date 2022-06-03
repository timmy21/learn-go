// 在 Go 语言中直接使用数组的情况并不多，更常用的数据结构是切片，即动态数组。
package main

import "fmt"

// https://github.com/golang/go/blob/master/src/runtime/slice.go#L15
// slice 运行时数据结构如下:
// type slice struct {
// 	array unsafe.Pointer // 指向底层数组
// 	len   int
// 	cap   int
// }

// append 元素时，如果 slice 容量不足以容纳新增元素时，会发生扩容。扩容规则如下:
//   1. 首先，如果新申请容量（cap）大于2倍的旧容量（old.cap），最终容量（newcap）就是新申请的容量（cap）
//   2. 否则，如果旧切片的容量小于256，则最终容量(newcap)就是旧容量(old.cap)的两倍，即（newcap=doublecap）
//   3. 否则，如果旧切片的容量大于等于256，则最终容量（newcap）从旧容量（old.cap）开始循环逐步增加容量，
//       即（newcap += (newcap + 3*threshold) / 4）直到最终容量（newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
//   4. 如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）
// 不同Go版本扩容规则可能会发生变化，具体请查看：https://github.com/golang/go/blob/master/src/runtime/slice.go#L178
func concat(a, b []int) []int {
	// slice 作为参数传递时，拷贝的是 "slice header"（大小为 3-word，64位机器是24字节），而不会拷贝底层的数组。
	return append(a, b...)
}

type Point struct {
	X int
	Y int
}

func (p Point) String() string {
	return fmt.Sprintf("{X:%d Y:%d}", p.X, p.Y)
}

// 由于 slice 的低成本拷贝以及灵活性，所以在 Go 编程时会被大量使用。而数组仅在一些特定场景使用。
// 关于 slice 的各种操作可以查看：https://github.com/golang/go/wiki/SliceTricks
func delete(a []*Point, i int) []*Point {
	copy(a[i:], a[i+1:])
	a[len(a)-1] = nil // or the zero value of T，释放内存
	return a[:len(a)-1]
}

func main() {
	// make([]T, length, capacity)
	s1 := make([]int, 0, 6)
	fmt.Println(s1)           // Output: []
	fmt.Println(s1[:cap(s1)]) // Output: [0 0 0 0 0 0]

	fmt.Println("==================")
	// 如果容量充足会复用底层数组
	s2 := concat(s1, []int{1, 2, 3})
	fmt.Println(s2)           // Output: [1 2 3]
	fmt.Println(cap(s2))      // Output: 6
	fmt.Println(s1)           // Output: []
	fmt.Println(s1[:cap(s1)]) // Output: [1 2 3 0 0 0]

	fmt.Println("==================")
	// 如果容量不足以容纳新增元素时，会发生扩容，创建新的底层数组
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
