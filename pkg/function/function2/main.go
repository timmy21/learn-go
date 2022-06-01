package main

import "fmt"

// 与其他赋值语句一样，函数参数使用“拷贝”方式传递的
func double1(nums [3]int) {
	for i := range nums {
		nums[i] *= 2
	}
}

// 仅拷贝值的“直接”部分，也就是通常说的“浅拷贝”
func double2(nums []int) {
	for i := range nums {
		nums[i] *= 2
	}
}

func main() {
	nums1 := [3]int{1, 2, 3}
	double1(nums1)
	fmt.Println(nums1) // output: [1 2 3]

	nums2 := []int{1, 2, 3}
	double2(nums2)
	fmt.Println(nums2) // output: [2 4 6]
}
