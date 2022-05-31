// Go 语言只有 for 一种循环结构，可以使用它实现各种循环需求
package main

import "fmt"

func sum1(n int) int {
	var s int
	// 完整的循环结构，包括: InitStmt, Condition，PostStmt 三部分
	for i := 1; i <= n; i++ {
		s += i
	}
	return s
}

func sum2(n int) int {
	var s int
	i := 1
	// 仅包含 Condition，类似于 while loop
	for i <= n {
		s += i
		i++
	}
	return s
}

func sum3(n int) int {
	var s int
	i := 1
	// 无限循环
	for {
		if i > n {
			break
		}
		s += i
		i++
	}
	return s
}

func main() {
	fmt.Println(sum1(5))
	fmt.Println(sum2(5))
	fmt.Println(sum3(5))
}
