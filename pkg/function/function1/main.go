package main

import "fmt"

func sum1(a, b, c int) int {
	return a + b + c
}

// 支持可变数量的参数
func sum2(nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

// 支持多值返回，支持命名返回值
func stats(first int, remain ...int) (min, max, sum int) {
	min = first
	max = first
	sum = first
	for _, num := range remain {
		sum += num
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return
}

func main() {
	fmt.Println(sum1(1, 2, 3))
	fmt.Println(sum2(1, 2, 3, 4))

	// 可以使用下划线忽略一个或多个返回值
	min, max, _ := stats(1, 2, 3)
	fmt.Printf("min: %d, max: %d\n", min, max)
}
