package main

import (
	"fmt"
)

// 在 Go1.18 之前，接口定义中只能包含“方法元素”，现在可以包含“类型元素”
// 在 Go1.18 之前，接口被看作定义了一个“方法集合”，一个类型实现了所有的这些方法被认为实现了这个接口
// 但在 Go1.18 之后，接口被视作“类型集合”，一个类型在这个接口的类型集合中，那么这个类型被认为实现了这个接口
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// 泛型函数包含“类型参数”，每一个类型参数都有一个“类型约束”，类型约束必须是接口
func Min[T Ordered](s []T) T {
	r := s[0] // 如果 s 为空，会产生 panic
	for _, v := range s[1:] {
		if v < r {
			r = v
		}
	}
	return r
}

// 下面等效于：Max[T interface{~int | ~float32 | ~float64}]
func Max[T ~int | ~float32 | ~float64](s []T) T {
	r := s[0] // 如果 s 为空，会产生
	for _, v := range s[1:] {
		if v > r {
			r = v
		}
	}
	return r
}

// 接口 ABInt 的类型集合可以认为是三个接口(Integer, A, B)类型集合的交集
type ABInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
	A()
	B()
}

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type A interface {
	A()
}

type B interface {
	B()
}

func main() {
	// 由于类型推断，调用泛型函数时，如果可以从函数参数中推断出类型参数时， 可以省略类型参数
	fmt.Println(Min([]int{1, 2, 3})) // 等价于：Min[int]([]int{1, 2, 3})
	fmt.Println(Max([]int{1, 2, 3}))
}
