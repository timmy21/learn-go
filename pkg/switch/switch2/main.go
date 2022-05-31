package main

import (
	"fmt"
	"time"
)

// 区别于某些语言需要显示调用 break，Go 默认会在 case 语句块的最后执行 break。
// 如果还需要执行下一个 case 语句块，可以调用 fallthrough。
func Greet(now time.Time) {
	switch hour := now.Hour(); { //  缺失的 switch 表达式等价于 "true"
	case hour < 12:
		fmt.Println("Good morning!")
		fallthrough
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}

func main() {
	Greet(time.Now())

	fmt.Println("============")
	Greet(time.Date(2022, 6, 1, 10, 0, 0, 0, time.Local))
}
