// 可以使用 range 进行迭代的数据结构包括：array, array指针, string, slice, map, chan
// range 不支持自定义数据类型。
package main

import (
	"fmt"
)

func sum1(nums []int) int {
	var s int
	for _, v := range nums {
		s += v
	}
	return s
}

// https://github.com/golang/gofrontend/blob/e387439bfd24d5e142874b8e68e7039f74c744d7/go/statements.cc#L5501
func sum2(nums []int) int {
	var s int

	// 内部实现大致等价于下面语句，
	{
		var index int
		var value int
		range_temp := nums // 待迭代的值会发生一次拷贝
		len_temp := len(range_temp)
		for index = 0; index < len_temp; index++ {
			value = range_temp[index]

			// 原始循环中的语句
			s += value
		}
	}
	return s
}

func sum3(nums []int) int {
	var s int

	// 通过 sum2 描述可以得出，其实在整个循环中使用的是同一个 v 变量
	// 这通常会被认为是一个 Go 语言的陷阱，就算知道这个陷阱并且经验丰富，但依然可能犯错。
	// 这个陷阱的场景两种情况：
	//   1. 取元素的地址，循环体外使用
	//   2. 通过闭包捕获元素，循环体外使用
	var copyed []*int
	for _, v := range nums {
		copyed = append(copyed, &v)
	}
	for _, v := range copyed {
		s += *v
	}
	return s
}

func sum4(nums []int) int {
	var s int

	var fns []func()
	for _, v := range nums {
		fns = append(fns, func() {
			s += v
		})
	}
	for _, fn := range fns {
		fn()
	}
	return s
}

func sum5(nums map[string]int) int {
	var s int
	// 在遍历 map 时，可以添加、删除 key。添加的 key 可能会在出现在接下来的迭代中，也可能不出现。
	for k, v := range nums {
		if v == 0 {
			delete(nums, k)
		}
		s += v
	}
	return s
}

func sum6(nums <-chan int) int {
	var s int
	for v := range nums {
		s += v
	}
	return s
}

func main() {
	fmt.Println(sum1([]int{1, 2, 3, 4, 5})) // Output: 15
	fmt.Println(sum2([]int{1, 2, 3, 4, 5})) // Output: 15
	fmt.Println(sum3([]int{1, 2, 3, 4, 5})) // Output: 25
	fmt.Println(sum4([]int{1, 2, 3, 4, 5})) // Output: 25

	fmt.Println("============")
	m := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}
	fmt.Println(sum5(m))
	fmt.Println(m)

	c := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		c <- i
	}
	close(c)
	fmt.Println(sum6(c))
}
